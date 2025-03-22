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
const (
	swahiliGreetingWithPrefix = "Jambo, "
	englishGreetingWithPrefix = "Hello, "
	spanishGreetingWithPrefix = "Hola, "
	frenchGreetingWithPrefix  = "Salut, "
	chineseGreetingWithPrefix = "Nǐ hǎo, "
)

func Jambo(Name, Language string) string {
	response := ""
	if Name == "" || Language == "" {
		response = swahiliGreetingWithPrefix + "wetu"
	}

	if Language == "Swahili" {
		response = swahiliGreetingWithPrefix + Name
	}

	if Language == "English" {
		response = englishGreetingWithPrefix + Name
	}

	if Language == "Spanish" {
		response = spanishGreetingWithPrefix + Name
	}

	if Language == "French" {
		response = frenchGreetingWithPrefix + Name
	}

	if Language == "Chinese" {
		response = chineseGreetingWithPrefix + Name
	}
	return response
}

func main() {
	fmt.Println(Jambo("Peter", "english"))
}
