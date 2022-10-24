package banking

import "fmt"

type Wallet struct{}

func (w Wallet) Deposit(money int) {
	fmt.Println(money)
}

func (w Wallet) Balance() int {
	return 0
}
