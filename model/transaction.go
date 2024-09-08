package model

type Transaction struct {
	ID uint64
	AccountIdFrom uint
	AccountIdTo uint
	Amount int
}
