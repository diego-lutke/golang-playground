// adapted from: https://withcodeexample.com/golang-tutorial-for-beginners/
package main

import "fmt"

func main() {
	// Basic for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// For loop with a condition
	// You can use a for loop with just a condition, similar to a while loop
	// in other languages.
	counter := 0
	for counter < 5 {
		fmt.Println(counter)
		counter++
	}

	// Infinite loop
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break // Exit the loop when i equals 5
		}
	}

	// For loop with range (used for iterating over slices, arrays, maps, etc.)
	// The range keyword is used in a for loop to iterate over elements in
	// slices, arrays, maps, and other data structures.
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Nested for loop
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}

	// Loop control statements: break and continue
	for i := 0; i < 5; i++ {
		if i == 2 {
			// Skip the current iteration when i equals 2
			continue
		}
		fmt.Println(i)
		if i == 3 {
			// Exit the loop when i equals 3
			break
		}
	}
}
