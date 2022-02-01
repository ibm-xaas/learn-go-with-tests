package main

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(deposit int) {
	w.balance += deposit
}

func (w *Wallet) Balance() int {
	return w.balance
}

func main() {

}
