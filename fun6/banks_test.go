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

	assertError := func(t testing.TB, got, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("Didn't get an error but wanted one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
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
		// no need to initial _ because it represents no value
		// Therefore = is sufficient instead of := notation
		_ = wallet.Withdraw(Bitcoin(10))

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
		assertError(t, err, ErrInsufficientFunds)
	})
}
