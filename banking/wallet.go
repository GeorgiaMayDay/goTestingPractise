package banking

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(money int) {
	w.balance += money
}

func (w *Wallet) Balance() int {
	return w.balance
}
