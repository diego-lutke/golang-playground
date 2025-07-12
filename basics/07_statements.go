// adapted from: https://withcodeexample.com/golang-tutorial-for-beginners/
package main

import "fmt"

func main() {
	// Simple if statement
	age := 25
	if age >= 18 {
		fmt.Println("You are eligible to vote.")
	}

	// If statement with a short statement
	// You can initialize a variable in the if statement itself, and its
	// scope will be limited to that block.
	if anotherAge := 25; anotherAge >= 18 {
		fmt.Println("You are eligible to vote.")
	}

	// If-else statement
	youngAge := 15
	if youngAge >= 18 {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote yet.")
	}

	// If-else if-else statement
	score := 85
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

	// Switch statement
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("It's the start of the week.")
	case "Friday":
		fmt.Println("It's almost the weekend.")
	default:
		fmt.Println("It's a regular day.")
	}
}
