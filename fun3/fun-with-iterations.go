// fun3: iteration
package fun3

import "strings"

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
`go test -bench=. -benchmen`
```text
10000000           25.70 ns/op           8 B/op           1 allocs/op
PASS

the `-benchmem` flag reports information about memory allocations:
- `B/op`: the number of bytes allocated per iteration
- `allocs/op`: the number of memory allocations per iteration
```
*/
func Repeat(letter string) string {
	var retLetter strings.Builder
	for range repeatedCount {
		// += called the Add AND assignment operator""
		// retLetter += letter
		retLetter.WriteString(letter)
	}
	return retLetter.String()
}

/*
Practice exercises

- Change the test so a caller can specify how many times the character is
  repeated and then fix the code
- Write ExampleRepeat to document your function
- Have a look through the strings package.
  Find functions you think could be useful and experiment with them by
  writing tests like we have here.
  Investing time learning the standard library will really pay off over time.
*/
/*
we are going to append all the code with duplicated function with a prefix EX
To do the practical exercises
*/
func ExRepeat(character string, n int) string {
	var repeatedString strings.Builder
	for range n {
		repeatedString.WriteString(character)
	}
	return repeatedString.String()
}

func LenRepeat(word string) int {
	lengthWord := len(word)
	return lengthWord
}

func ExpiRepeat(character string, n int) string {
	return strings.Repeat(character, n)
}
