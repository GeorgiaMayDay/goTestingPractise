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

	t.Run("withdraw, insufficient funds", func(t *testing.T) {
		err := wallet.Withdraw(Bitcoin(100))
		want := Bitcoin(5)

		assertBalance(t, wallet, want)

		if err == nil {
			t.Errorf("wanted error, got nothing")
		}

	})

}
