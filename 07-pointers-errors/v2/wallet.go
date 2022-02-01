package main

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(deposit Bitcoin) {
	w.balance += deposit
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(withdraw Bitcoin) error {

	if withdraw > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= withdraw
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {

}
