package storage

import (
	"context"

	"app/api/models"
)

type StorageI interface {
	Close()
	User() UserRepoI
	Phone() PhoneRepoI
}

type UserRepoI interface {
	Create(context.Context, *models.UserCreate) (string, error)
	GetByID(context.Context, *models.UserPrimaryKey) (*models.User, error)
	GetList(context.Context, *models.UserGetListRequest) (*models.UserGetListResponse, error)
	Update(context.Context, *models.UserUpdate) (int64, error)
	Delete(context.Context, *models.UserPrimaryKey) error
}

type PhoneRepoI interface {
	Create(context.Context, *models.PhoneCreate) (string, error)
	GetByID(context.Context, *models.PhonePrimaryKey) (*models.Phone, error)
	GetList(context.Context, *models.PhoneGetListRequest) (*models.PhoneGetListResponse, error)
	Update(context.Context, *models.PhoneUpdate) (int64, error)
	Delete(context.Context, *models.PhonePrimaryKey) error
}
