package repository

import "gopher-bank/model"

type UserRepository interface {
	Save(user *model.User) (*model.User, error)
	FindAll() ([]model.User, error)
}
