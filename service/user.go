package service

import (
	"clean-architecture/model"
	"clean-architecture/repo"
	"context"
)

type User struct {
	repo repo.IRepo
}

func NewUser(repo repo.IRepo) *User {
	return &User{
		repo: repo,
	}
}

type IUser interface {
	GetUsers(ctx context.Context) (users []model.User, err error)
}

func (s *User) GetUsers(ctx context.Context) (users []model.User, err error) {
	// Get user from repo
	users, err = s.repo.GetUsers(ctx)
	if err != nil {
		return users, err
	}
	return users, nil
}
