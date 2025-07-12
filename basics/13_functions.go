// adapted from: https://withcodeexample.com/golang-tutorial-for-beginners/
package main

import "fmt"

// Function named 'sayHello'
func sayHello() {
	fmt.Println("Hello, Go!")
}

// Function with parameters
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Function with multiple parameters
func add(a, b int) {
	result := a + b
	fmt.Printf("%d + %d = %d\n", a, b, result)
}

// Function with a return value
func multiply(a, b int) int {
	return a * b
}

// Function with named return values
func divide(a, b float64) (quotient float64, remainder float64) {
	quotient = a / b
	remainder = a - (quotient * b)
	return
}

// Variadic function (accepts a variable number of arguments)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function that takes another function as a parameter
// One function can pass another function as a parameter.
// A operate function that accepts two numbers and an additional
// function as parameters is demonstrated in this example.
func operate(a, b int, operation func(int, int) int) {
	result := operation(a, b)
	fmt.Println("Result:", result)
}

func main() {
	// Calling the 'sayHello' function
	sayHello()

	// Calling the 'greet' function with an argument
	greet("Alice")
	greet("Bob")

	// Calling the 'add' function with two arguments
	add(3, 7)

	// Calling the 'multiply' function and using the returned value
	result := multiply(4, 5)
	fmt.Println("Result:", result)

	// Calling the 'divide' function and using the returned values
	q, r := divide(10, 3)
	fmt.Printf("Quotient: %.2f, Remainder: %.2f\n", q, r)

	// Calling the 'sum' function with a variable number of arguments
	sumResult := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", sumResult)

	// Anonymous function (closure)
	anonymousAdd := func(a, b int) int {
		return a + b
	}
	// Calling the anonymous function
	anonymousResult := anonymousAdd(8, 12)
	fmt.Println("Result:", anonymousResult)

	// Calling the 'operate' function with the 'anonymousAdd'
	// function as a parameter
	operate(10, 5, anonymousAdd)
}
