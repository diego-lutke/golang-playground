// adapted from: https://withcodeexample.com/golang-tutorial-for-beginners/
package basics

import "fmt"

func Arrays() {
	// Declaration and initialization of an array
	var numbers [5]int                    // An array of integers with a size of 5
	numbers = [5]int{1, 2, 3, 4, 5}       // Initializing array elements individually
	moreNumbers := [5]int{6, 7, 8, 9, 10} // Short declaration with initialization
	fmt.Println("Numbers:", numbers)
	fmt.Println("More Numbers:", moreNumbers)

	// Accessing array elements
	// obs: the index starts from 0.
	firstElement := numbers[0]
	thirdElement := numbers[2]
	fmt.Println("First Element:", firstElement)
	fmt.Println("Third Element:", thirdElement)

	// Getting the length of the array
	length := len(numbers)
	fmt.Println("Length of the array:", length)

	// Iterating over array elements
	// The `range` keyword is used to iterate over the elements of an array,
	// providing both the index and value at each iteration.
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Declaration and initialization of a 2D array
	var matrix [3][3]int
	matrix = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix:", matrix)

	// Declaration and creation of a slice

	// Once an array has allocated its size, the size can no longer be changed.
	// A slice, on the other hand, is a variable length version of an array

	// The size of a slice is not specified in the declaration, and it is
	// determined by the number of elements provided during initialization.
	var numbersForSlice []int                    // Declaration of an integer slice
	numbersForSlice = []int{1, 2, 3, 4, 5}       // Initialization of the slice
	moreNumbersForSlice := []int{6, 7, 8, 9, 10} // Short declaration with initialization

	fmt.Println("Numbers:", numbersForSlice)
	fmt.Println("More Numbers:", moreNumbersForSlice)

	// Accessing and modifying slice elements
	firstSliceElement := numbersForSlice[0]
	// thirdSliceElement := numbersForSlice[2]
	fmt.Println("First Element:", firstSliceElement)
	numbersForSlice[1] = 20 // Modifying the second element
	fmt.Println("Modified Slice:", numbersForSlice)
	numbersForSlice[1] = 2

	// Slicing a slice
	slicedNumbers := numbersForSlice[1:4]
	fmt.Println("Sliced Numbers:", slicedNumbers)

	// Appending elements to a slice
	numbersForSlice = append(numbersForSlice, 6, 7, 8)
	fmt.Println("Updated Slice:", numbersForSlice)

	// Length of the slice
	// The len() function returns the length (number of elements) of a slice
	numbersLength := len(numbersForSlice)

	// Capacity of the slice
	// the cap() function returns the capacity (maximum number of elements it
	// can hold without resizing).
	capacity := cap(numbersForSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", numbersLength, capacity)

	// Copying elements from one slice to another
	numbersToCopy := []int{1, 2, 3, 4, 5}
	anotherNumbers := make([]int, len(numbersToCopy))
	copy(anotherNumbers, numbersToCopy)
	fmt.Println("Copied Slice:", anotherNumbers)
}
