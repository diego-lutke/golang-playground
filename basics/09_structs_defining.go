// adapted from: https://withcodeexample.com/golang-tutorial-for-beginners/
package basics

import "fmt"

// A struct is a composite data type in Go that unifies variables, also known
// as fields, under a single name. A struct’s fields may each have a distinct
// data type. Using structures, you may design user-defined types that group
// and encapsulate similar data points.

// Defining a struct named 'Person' with three fields: FirstName, LastName,
// and Age
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func StructsDefining() {
	// Creating an instance of the 'Person' struct
	person1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// Accessing struct fields using dot notation
	fmt.Println("First Name:", person1.FirstName)
	fmt.Println("Last Name:", person1.LastName)
	fmt.Println("Age:", person1.Age)
}
