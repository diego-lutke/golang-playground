// adapted from: https://withcodeexample.com/golang-tutorial-for-beginners/
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

	// String Formatting
	name := "Alice"
	age := 25

	// The fmt.Sprintf function is used for string formatting.
	// It returns a formatted string without printing it to the console.
	message := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(message)

	// Go has excellent support for Unicode characters in strings,
	// allowing you to work with a wide range of international characters.
	// The Japanese phrase “Hello, World!” is represented by Unicode
	// characters in the given string.
	unicodeString := "こんにちは, 世界!" // Hello, World! in Japanese
	fmt.Println(unicodeString)
}
