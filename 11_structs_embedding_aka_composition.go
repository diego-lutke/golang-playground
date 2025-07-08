// adapted from: https://withcodeexample.com/golang-tutorial-for-beginners/
package main

import "fmt"

// Struct embedding allows one struct to be embedded within another,
// forming a composition relationship.

// Defining a 'Contact' struct
type Contact struct {
	Email string
	Phone string
}

// Defining a 'Client' struct with embedded 'Contact'
type Client struct {
	FirstName string
	LastName  string
	Age       int
	Contact   // Embedding the 'Contact' struct
}

func main() {
	// Creating an instance of the 'Client' struct with embedded 'Contact'
	client := Client{
		FirstName: "Bob",
		LastName:  "Johnson",
		Age:       35,
		Contact: Contact{
			Email: "bob@example.com",
			Phone: "555-1234",
		},
	}

	fmt.Println("First Name:", client.FirstName)
	fmt.Println("Last Name:", client.LastName)
	fmt.Println("Age:", client.Age)
	fmt.Println("Email:", client.Email) // Accessing embedded struct field
	fmt.Println("Phone:", client.Phone) // Accessing embedded struct field
}
