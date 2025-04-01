package fun6

import "fmt"

/*
Since we are creating a bitcoin,
let's be descriptive with a new type from existing ones
*/

type (
	Bitcoin int
	Wallet  struct {
		balance Bitcoin
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
