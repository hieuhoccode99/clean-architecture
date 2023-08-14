package repo

import (
	"clean-architecture/model"
	"context"
)

func (r *Repo) GetUsers(ctx context.Context) ([]model.User, error) {
	// Lấy từ database
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return []model.User{}, err
	}
	return users, nil
}
