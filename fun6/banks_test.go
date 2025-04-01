package fun6

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	// print the address of the memory
	fmt.Printf("address of balance in test is %p \n", &wallet.balance)
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
