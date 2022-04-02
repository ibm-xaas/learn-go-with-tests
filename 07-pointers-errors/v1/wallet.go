package main

// Wallet ...
type Wallet struct {
	balance int
}

// Deposit ...
func (w *Wallet) Deposit(deposit int) {
	w.balance += deposit
}

// Balance ...
func (w *Wallet) Balance() int {
	return w.balance
}

func main() {

}
