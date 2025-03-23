package fun3

/*
Write the minimal amout of code for the test to run and check the failing test output
All you need to go right now is enough to make it compile so you can check your test is written well.

For example: this example is enough to compile and verify it your testing case works.
```go
package iteration

	func Repeat(character string) string{
	   return ""
	}

```
*/
/*
Refactor1: create a variable for the number of times to repeat a character
					 since it won't change in throughout lifetime execution of the program
					 we will use a constant
*/
const repeatedCount = 7

/*
In golang, strings are immutable therefore everytime we append a character, a new memory slot
is created, the process to clean up the old one, it tiggers at the same time.
Hence, the code is not optimize performance wise.
Go provides us with the method strings.Builder of type stringsBuilder which minimize memory
copiying and to get the string you use the method String
*/
func Repeat(letter string) string {
	retLetter := ""
	for range repeatedCount {
		// += called the Add AND assignment operator""
		retLetter += letter
	}
	return retLetter
}
