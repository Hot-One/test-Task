package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"app/api/models"
	"app/pkg/helper"
)

// Login godoc
// @ID login
// @Router /user/login [POST]
// @Summary Login
// @Description Login
// @Tags Login
// @Accept json
// @Procedure json
// @Param login body models.UserLogin true "LoginRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) Login(c *gin.Context) {
	var login models.UserLogin

	err := c.ShouldBindJSON(&login) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create user", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Login: login.Login})
	if err != nil {
		if err.Error() == "no rows in result set" {
			h.handlerResponse(c, "User does not exist", http.StatusInternalServerError, nil)
			return
		}
		h.handlerResponse(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(login.Password))
	if err != nil {
		h.handlerResponse(c, "bcrypt.CompareHashAndPassword", http.StatusBadRequest, err.Error())
		return
	}

	// if resp.Password != login.Password {
	// 	h.handlerResponse(c, "Wrong password", http.StatusInternalServerError, nil)
	// 	return
	// }
	token, err := helper.GenerateJWT(map[string]interface{}{
		"user_id":    resp.Id,
		"user_name":  resp.Name,
		"user_age":   resp.Age,
		"user_login": resp.Login,
	}, time.Hour*360, h.cfg.SecretKey)
	if err != nil {
		h.handlerResponse(c, "Error while Generate JWT", http.StatusInternalServerError, nil)
		return
	}

	c.SetCookie("user", token, 3600, "/", "localhost", false, true)

	h.handlerResponse(c, "token", http.StatusCreated, token)
}

// Register godoc
// @ID register
// @Router /user/register [POST]
// @Summary Register
// @Description Register
// @Tags Register
// @Accept json
// @Procedure json
// @Param register body models.UserCreate true "CreateUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) Register(c *gin.Context) {
	var createUser models.UserCreate
	var id string

	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		h.handlerResponse(c, "error user should bind json", http.StatusBadRequest, err.Error())
		return
	}

	if len(createUser.Password) < 7 {
		h.handlerResponse(c, "Password should inculude more than 7 elements", http.StatusBadRequest, errors.New("Password len should inculude more than 8 elements"))
		return
	}

	bcryprP, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), bcrypt.DefaultCost)
	if err != nil {
		h.handlerResponse(c, "bcrypt.GenerateFromPassword", http.StatusInternalServerError, err.Error())
		return
	}

	createUser.Password = string(bcryprP)

	resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Login: createUser.Login})
	if err != nil {
		if err.Error() == "no rows in result set" {
			id, err = h.strg.User().Create(context.Background(), &createUser)
			if err != nil {
				h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			h.handlerResponse(c, "User already exist", http.StatusInternalServerError, err.Error())
			return
		}
	} else if err == nil {
		h.handlerResponse(c, "User already exist", http.StatusBadRequest, nil)
		return
	}
	resp, err = h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "Register=> h.strg.User().GetByID", http.StatusInternalServerError, nil)
	}

	h.handlerResponse(c, "create user resposne", http.StatusCreated, resp)
}
