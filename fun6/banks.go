package fun6

type Wallet struct {
	balance int
}

/*
func (w Wallet) Deposit(amount int)
In go when you call a function or a method the arguments are copied
*/
func (w Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
