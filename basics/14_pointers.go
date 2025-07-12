// adapted from: https://withcodeexample.com/golang-tutorial-for-beginners/
package main

import "fmt"

// Pointers are variables in Go that hold the address of another variable
// in memory. They serve as an oblique reference to the value kept in a
// certain memory region. This is a summary of how to use pointers in Go:

// Defining a struct
type Toy struct {
	Name   string
	Type   string
	Height float32
}

func main() {
	// Declaring a variable and a pointer
	var value int = 42
	var pointer *int

	// Initializing the pointer with the address of the variable
	pointer = &value

	// Dereferencing the pointer to get the value it points to
	// The * operator is used to dereference a pointer, which means
	// accessing the value stored at the memory address held by the pointer.
	dereferencedValue := *pointer

	fmt.Println("Value:", value)
	fmt.Println("Pointer:", pointer)
	fmt.Println("Dereferenced Value:", dereferencedValue)

	// Modifying the variable through the pointer
	// Since the pointer holds the memory address of a variable, modifying
	// the value through the pointer also modifies the original variable.
	*pointer = 100
	fmt.Println("Modified Value:", value)

	// Declaring a nil pointer
	// A pointer that is not assigned any memory address is called a “nil”
	// pointer. It is the zero value for pointers in Go.
	var anotherPointer *int
	// Checking if the pointer is nil
	if anotherPointer == nil {
		fmt.Println("Pointer is nil")
	} else {
		fmt.Println("Pointer is not nil")
	}

	// Using the 'new' function to create a pointer
	// The new function is used to create a new variable and return
	// its memory address as a pointer.
	anotherPointerAgain := new(int)
	// Assigning a value through the pointer
	*anotherPointerAgain = 42
	fmt.Println("Value assigned through 'new' pointer:", *anotherPointerAgain)

	// A struct’s fields can be accessed and changed via pointers.
	// In this example, the fields of a `Toy` struct are accessed
	// by creating a reference to the struct.
	// Creating an instance of the struct
	toy := Toy{
		Name:   "Soldier",
		Type:   "Figurine",
		Height: 10.65,
	}
	// Creating a pointer to the struct
	var toyPointer *Toy
	toyPointer = &toy
	// Accessing struct fields through the pointer
	fmt.Println("Name:", toyPointer.Name)
	fmt.Println("Type:", toyPointer.Type)
	fmt.Println("Height:", toyPointer.Height)
}
