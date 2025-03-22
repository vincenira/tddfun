package main

import "fmt"

/*
Here we separated the domain "code" from the side effect
which is the print of the word jambo wetu
because the print and instantiation for the word allows
an easy refactoring and separated concern
*/

// Const information on the variable helps to capture the essential of the variable and also aid the performance of the program
// the essential, here is that the variable will not change throughout the execution life of the program
const greetingWithPrefix = "Jambo, "

func Jambo(Name string) string {
	return greetingWithPrefix + Name
}

func main() {
	fmt.Println(Jambo("Peter"))
}
