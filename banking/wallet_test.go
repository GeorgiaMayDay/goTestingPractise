package banking

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	t.Run("deposit", func(t *testing.T) {
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("deposit", func(t *testing.T) {
		wallet.Withdraw(5)
		want := Bitcoin(5)

		assertBalance(t, wallet, want)
	})

}
