package service

import (
	"errors"
	"gopher-bank/model"
	"gopher-bank/repository"
	"math/rand"
)

var (
	userRepository repository.UserRepository
)

type service struct{}

type UserService interface {
	Validate(user *model.User) error
	Create(user *model.User) (*model.User, error)
	FindAll() ([]model.User, error)
}

func NewUserService(repository repository.UserRepository) UserService {
	userRepository = repository
	return &service{}
}

func (s *service) Validate(user *model.User) error {
	if user == nil {
		err := errors.New("the user is empty")
		return err
	}
	if user.Username == "" {
		err := errors.New("the user username is empty")
		return err
	}
	if user.Email == "" {
		err := errors.New("the user email is empty")
		return err
	}
	if user.Password == "" {
		err := errors.New("the user password is empty")
		return err
	}
	return nil
}

func (s *service) Create(user *model.User) (*model.User, error) {
	user.ID = rand.Uint64()
	return userRepository.Save(user)
}

func (s *service) FindAll() ([]model.User, error) {
	return userRepository.FindAll()
}
