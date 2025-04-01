package fun6

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		// print the address of the memory
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		want := Bitcoin(10)

		// Update our test format to use strings
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		// print the address of the memory
		want := Bitcoin(10)

		// Update our test format to use strings
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
