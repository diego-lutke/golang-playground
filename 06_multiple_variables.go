// code source: https://withcodeexample.com/golang-tutorial-for-beginners/
package main

import "fmt"

func main() {
	// Multiple variables
	// You can declare multiple variables at once using the var keyword
	// with parentheses.
	var (
		firstName = "Alice"
		lastName  = "Smith"
		age       = 30
	)

	fmt.Printf("Name: %s %s, Age: %d\n", firstName, lastName, age)
}
