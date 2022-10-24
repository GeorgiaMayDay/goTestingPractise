package banking

import (
	"testing"
)

func TestWallet(t *testing.T){
	wallet := Wallet{}

	wallet.Deposit(5)

	got := wallet.Balance()
	want := 5

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}


}