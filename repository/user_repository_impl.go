package repository

import "gopher-bank/model"

var (
	users []model.User
)

type repository struct{}

func NewUserRepository() UserRepository {
	return &repository{}
}

func (r *repository) FindAll() ([]model.User, error) {
	return users, nil
}

func (r *repository) Save(user *model.User) (*model.User, error) {
	users = append(users, *user)
	return user, nil
}
