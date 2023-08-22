package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageI, logger logger.LoggerI) {

	handler := handler.NewHandler(cfg, storage, logger)

	// User Api
	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.GetByIdUser)
	r.GET("/user", handler.GetListUser)
	r.PUT("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
