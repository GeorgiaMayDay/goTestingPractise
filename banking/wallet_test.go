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

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	t.Run("deposit", func(t *testing.T) {
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet.Withdraw(Bitcoin(5))
		want := Bitcoin(5)

		assertBalance(t, wallet, want)

	})

	t.Run("withdraw, insufficient funds", func(t *testing.T) {
		err := wallet.Withdraw(Bitcoin(100))
		want := Bitcoin(5)

		assertBalance(t, wallet, want)

		assertError(t, err, "Insufficient funds, can't withdraw")

	})

}
