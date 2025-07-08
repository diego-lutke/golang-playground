// adapted from: https://withcodeexample.com/golang-tutorial-for-beginners/
package main

import "fmt"

// Defining a 'Human' struct
type Human struct {
	FirstName string
	LastName  string
	Age       int
}

// Adding a method to the 'Human' struct
func (p Human) FullName() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	// Creating an instance of the 'Human' struct
	human := Human{
		FirstName: "Charlie",
		LastName:  "Brown",
		Age:       28,
	}

	// Calling the 'FullName' method
	fullName := human.FullName()

	fmt.Println("Full Name:", fullName)
}
