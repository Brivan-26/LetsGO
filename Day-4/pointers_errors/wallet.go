package pointers_errors

import (
	"errors"
)

var ErrInsufficientFunds  = errors.New("Not sufficient amount")

type Ether int

type Wallet struct {
	balance Ether
}

func (w *Wallet) Deposit(amount Ether) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Ether) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Ether {
	return w.balance
}