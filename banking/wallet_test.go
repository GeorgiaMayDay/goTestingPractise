package banking

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	t.Run("deposit", func(t *testing.T) {
		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("deposit", func(t *testing.T) {
		wallet.Withdraw(5)

		got := wallet.Balance()
		want := Bitcoin(5)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}
