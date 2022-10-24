package banking

import "fmt"

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(money Bitcoin) {
	w.balance += money
}

func (w *Wallet) Withdraw(money Bitcoin) error {
	w.balance -= money
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
