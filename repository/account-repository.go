package repository

import "gopher-bank/model"

type AccountRepository interface {
	Save(account *model.Account) (*model.Account, error)
	FindById(id uint64) (*model.Account, error)
}
