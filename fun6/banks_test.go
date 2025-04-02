package fun6

import (
	"fmt"
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
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		// print the address of the memory
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		want := Bitcoin(10)

		// Update our test format to use strings
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		// print the address of the memory
		want := Bitcoin(10)

		// Update our test format to use strings
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		// If you see a function that takes
		// arguments or returns values that are interfaces, they can be nillable.
		if err == nil {
			t.Error("Wanted an error but didn't get one")
		}
	})
}
