package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"app/api/models"
	"app/pkg/helper"
)

type PhoneRepo struct {
	db *pgxpool.Pool
}

func NewPhoneRepo(db *pgxpool.Pool) *PhoneRepo {
	return &PhoneRepo{
		db: db,
	}
}

func (r *PhoneRepo) Create(ctx context.Context, req *models.UserCreate) (string, error) {
	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO users(id, login, password, name, age)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(ctx, query,
		id,
		req.Login,
		req.Password,
		req.Name,
		req.Age,
	)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return id, nil
}

func (r *PhoneRepo) GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {

	if len(req.Login) > 0 {
		var (
			query string

			id       sql.NullString
			login    sql.NullString
			password sql.NullString
			name     sql.NullString
			age      sql.NullInt64
		)
		query = `
		SELECT
			id,
			login,
			password,
			name,
			age
		FROM users
		WHERE login = $1
	`

		err := r.db.QueryRow(ctx, query, req.Login).Scan(
			&id,
			&login,
			&password,
			&name,
			&age,
		)

		if err != nil {
			return nil, err
		}

		return &models.User{
			Id:       id.String,
			Login:    login.String,
			Password: password.String,
			Name:     name.String,
			Age:      int(age.Int64),
		}, nil
	}

	var (
		query string

		id       sql.NullString
		login    sql.NullString
		password sql.NullString
		name     sql.NullString
		age      sql.NullInt64
	)

	query = `
		SELECT
			id,
			login,
			password,
			name,
			age
		FROM users
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&login,
		&password,
		&name,
		&age,
	)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:       id.String,
		Login:    login.String,
		Password: password.String,
		Name:     name.String,
		Age:      int(age.Int64),
	}, nil
}

func (r *PhoneRepo) GetList(ctx context.Context, req *models.UserGetListRequest) (*models.UserGetListResponse, error) {

	var (
		resp   = &models.UserGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			login,
			password,
			name,
			age
		FROM users
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND name ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id       sql.NullString
			login    sql.NullString
			password sql.NullString
			name     sql.NullString
			age      sql.NullInt64
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&login,
			&password,
			&name,
			&age,
		)

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &models.User{
			Id:       id.String,
			Login:    login.String,
			Password: password.String,
			Name:     name.String,
			Age:      int(age.Int64),
		})
	}

	return resp, nil
}

func (r *PhoneRepo) Update(ctx context.Context, req *models.UserUpdate) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)
	query = `
		UPDATE
			users
		SET
			id = :id,
			login = :login,
			password = :password,
			name = :name,
			age = :age
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":       req.Id,
		"login":    req.Login,
		"password": req.Password,
		"name":     req.Name,
		"age":      req.Age,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *PhoneRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM users WHERE id = $1", req.Id)

	if err != nil {
		return err
	}

	return nil
}
