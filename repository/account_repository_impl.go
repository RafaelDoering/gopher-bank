package repository

import (
	"gopher-bank/model"
	"slices"
)

var (
	accounts []model.Account
)

type accountRepository struct{}

func NewAccountRepository() AccountRepository {
	return &accountRepository{}
}

func (r *accountRepository) FindById(id uint64) (*model.Account, error) {
	idx := slices.IndexFunc(accounts, func(a model.Account) bool { return a.ID == id })


	return &accounts[idx], nil
}

func (r *accountRepository) Save(user *model.Account) (*model.Account, error) {
	accounts = append(accounts, *user)
	return user, nil
}
