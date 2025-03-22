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
	langSwa = "Swahili"
	langEng = "English"
	langSpa = "Spanish"
	langFre = "French"
	langChi = "Chinese"

	swahiliGreetingWithPrefix = "Jambo, "
	englishGreetingWithPrefix = "Hello, "
	spanishGreetingWithPrefix = "Hola, "
	frenchGreetingWithPrefix  = "Salut, "
	chineseGreetingWithPrefix = "Nǐ hǎo, "
)

func Jambo(Name, Language string) string {
	response := ""
	// if you find that there is a lot of if statement of a same variable with multiple values
	// it means a switch statement should replace it

	switch Language {
	case langSwa:
		response = swahiliGreetingWithPrefix + Name
	case langEng:
		response = englishGreetingWithPrefix + Name
	case langSpa:
		response = spanishGreetingWithPrefix + Name
	case langFre:
		response = frenchGreetingWithPrefix + Name
	case langChi:
		response = chineseGreetingWithPrefix + Name
	default:
		response = swahiliGreetingWithPrefix + "wetu"
	}

	return response
}

func main() {
	fmt.Println(Jambo("Peter", "english"))
}
