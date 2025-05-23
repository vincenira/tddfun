package fun6

import (
	"errors"
	"fmt"
)

/*
Since we are creating a bitcoin,
let's be descriptive with a new type from existing ones
*/

type (
	Bitcoin int
	Wallet  struct {
		balance Bitcoin
	}
	// This interface is defined in the fmt package and lets you define how your type is printed
	// when used with %s format string in prints
	Stringer interface {
		String() string
	}
)

/*
func (w Wallet) Deposit(amount int)
In go when you call a function or a method the arguments are copied
*/
/*
func (w Wallet) Deposit(amount int) {
	// print the address of balance in the method
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
*/
func (w *Wallet) Deposit(amount Bitcoin) {
	// print the address of balance in the method
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// As you can see, the syntax for creating a method on a type declaration is the same as it is on a struct.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// To signal a problem in go, it is idiomatic to return an error for the caller to check and act on

// Refactor error by being descriptive with our error
// var keyword allows us to define values global to the the package
var ErrInsufficientFunds = errors.New("Cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
