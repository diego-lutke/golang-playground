// code source: https://withcodeexample.com/golang-tutorial-for-beginners/
package main

import "fmt"

func main() {
	// Go handles strings as immutable, which means that once a string is formed,
	// it is impossible to modify its contents.

	// Declaring and initializing a string
	greeting := "Hello, Golang!"
	fmt.Println(greeting)

	// Calculating the lenght of a string
	// note the difference between Println and Printf
	length := len(greeting)
	fmt.Printf("Length of the string: %d\n", length)

	// Concatenating strings
	firstName := "John"
	lastName := "Doe"

	fullName := firstName + " " + lastName

	fmt.Println("Full Name:", fullName)

	// Indexing and Slicing
	// In Go, strings are made up of bytes. You can access individual characters
	// by indexing the string.
	stringForSlicing := "Hello, Gophers!"

	// Accessing individual characters (byte-wise)
	firstChar := stringForSlicing[0]

	fmt.Println("First Character:", string(firstChar))

	// Slicing a string
	slicedString := stringForSlicing[7:15]

	fmt.Println("Sliced String:", slicedString)
}
