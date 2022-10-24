package banking

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(5))

	got := wallet.Balance()
	want := Bitcoin(5)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
