// adapted from: https://withcodeexample.com/golang-tutorial-for-beginners/
package main

import "fmt"

func main() {
	// Creating an anonymous struct
	// Anonymous structs are structs defined without a specified name.
	// They are useful for creating one-time-use data structures.
	person := struct {
		FirstName string
		LastName  string
		Age       int
	}{
		FirstName: "Alice",
		LastName:  "Smith",
		Age:       25,
	}

	fmt.Println("First Name:", person.FirstName)
	fmt.Println("Last Name:", person.LastName)
	fmt.Println("Age:", person.Age)
}
