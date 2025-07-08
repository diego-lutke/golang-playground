// adapted from: https://withcodeexample.com/golang-tutorial-for-beginners/
package main

import "fmt"

func main() {
	// Implicit type declaration:
	// In this example, the type of the variable message is implicitly
	// inferred from the assigned value, which is a string.
	var message = "Hello, Go!"
	fmt.Println(message)

	// Explicit type declaration:
	// Here, we declare a variable age with an explicit type of int.
	// We later assign a value of 25 to it and print the result.
	var age int
	age = 25
	fmt.Println("Age:", age)

	// Short variable declaration:
	// The short variable declaration := is a concise way to declare and initialize
	// a variable. The type is inferred from the assigned value.
	name := "John"
	fmt.Println("Name:", name)

	// Multiple variables
	// You can declare multiple variables at once using the var keyword
	// with parentheses.
	var (
		firstName = "Alice"
		lastName  = "Smith"
		birthYear = 1992
	)
	fmt.Printf("Name: %s %s, Birth Year: %d\n", firstName, lastName, birthYear)

	// Constants
	// Constants are declared using the const keyword.
	// They must be assigned a value at the time of declaration
	// and cannot be changed later.
	const pi = 3.14159
	fmt.Println("Value of pi:", pi)

	// Zero values
	// A variable’s type determines the value allocated to it when it is
	// defined but not explicitly initialised. These examples address a variety
	// of topics related to using variables in Go. You will have a better
	// grasp of how variables work in Go programmes by experimenting
	// with these ideas.
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool
	var zeroArray [5]int
	var zeroArrayMatrix [3][3]int
	var zeroSlice []int
	var zeroMap map[string]string
	var zeroPointer *int
	fmt.Printf("Zero int: %d\n", zeroInt)
	fmt.Printf("Zero float: %f\n", zeroFloat)
	fmt.Printf("Zero string: %s\n", zeroString)
	fmt.Printf("Zero bool: %t\n", zeroBool)
	fmt.Println("Zero array:", zeroArray)
	fmt.Println("Zero array (2d):", zeroArrayMatrix)
	fmt.Println("Zero slice:", zeroSlice)
	fmt.Println("Zero map:", zeroMap)
	fmt.Println("Zero pointer:", zeroPointer)
}
