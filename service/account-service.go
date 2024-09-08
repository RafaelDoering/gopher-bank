package service

import (
	"errors"
	"gopher-bank/model"
	"gopher-bank/repository"
	"math/rand"
)

var (
	accountRepository repository.AccountRepository
)

type accountService struct{}

type AccountService interface {
	Validate(user *model.Account) error
	Create(user *model.Account) (*model.Account, error)
	FindById(id uint64) (*model.Account, error)
}

func NewAccountService(repository repository.AccountRepository) AccountService {
	accountRepository = repository
	return &accountService{}
}

func (s *accountService) Validate(account *model.Account) error {
	if account == nil {
		err := errors.New("the account is empty")
		return err
	}
	if account.UserId == 0 {
		err := errors.New("the account userId is empty")
		return err
	}
	return nil
}

func (s *accountService) Create(account *model.Account) (*model.Account, error) {
	account.ID = rand.Uint64()
	return accountRepository.Save(account)
}

func (s *accountService) FindById(id uint64) (*model.Account, error) {
	return accountRepository.FindById(id)
}
