package api

import "errors"

var (
	ErrInvalidParameter    = errors.New("invalid parameter")
	ErrSomethingWentWrong  = errors.New("something went wrong")
	ErrFailedCreatedWallet = errors.New("failed to create wallet")
	ErrNotEnoughMoney      = errors.New("not enough money on the wallet")
)
