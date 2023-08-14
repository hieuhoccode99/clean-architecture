package repo

import (
	"clean-architecture/model"
	"context"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}

type IRepo interface {
	GetUsers(ctx context.Context) ([]model.User, error)
}
