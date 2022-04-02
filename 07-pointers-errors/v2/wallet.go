package main

import (
	"errors"
	"fmt"
)

// Stringer ...
type Stringer interface {
	String() string
}

// Bitcoin ...
type Bitcoin int

// String ...
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet ...
type Wallet struct {
	balance Bitcoin
}

// Deposit ...
func (w *Wallet) Deposit(deposit Bitcoin) {
	w.balance += deposit
}

// ErrInsufficientFunds ...
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw ...
func (w *Wallet) Withdraw(withdraw Bitcoin) error {

	if withdraw > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= withdraw
	return nil
}

// Balance ...
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {

}
