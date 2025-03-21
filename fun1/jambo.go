package main

import "fmt"

/*
Here we separated the domain "code" from the side effect
which is the print of the word jambo wetu
because the print and instantiation for the word allows
an easy refactoring and separated concern
*/

func Jambo() string {
	return "Jambo wetu"
}

func main() {
	fmt.Println(Jambo())
}
