// adapted from: https://withcodeexample.com/golang-tutorial-for-beginners/
package basics

import "fmt"

func Operators() {
	// Arithmetic operators
	a := 10
	b := 3
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b)
	fmt.Println("Remainder:", a%b)
	fmt.Println("Increment:", a+1)
	fmt.Println("Decrement:", a-1)

	// Relational operators
	x := 5
	y := 8
	fmt.Println("Equal:", x == y)
	fmt.Println("Not Equal:", x != y)
	fmt.Println("Greater Than:", x > y)
	fmt.Println("Less Than:", x < y)
	fmt.Println("Greater Than or Equal To:", x >= y)
	fmt.Println("Less Than or Equal To:", x <= y)

	// Logical operators
	isTrue := true
	isFalse := false
	fmt.Println("AND:", isTrue && isFalse)
	fmt.Println("OR:", isTrue || isFalse)
	fmt.Println("NOT:", !isTrue)

	// Assignment operators
	c := 10
	d := 5
	c += d // Equivalent to c = c + d
	fmt.Println("c:", c)
	c -= d // Equivalent to c = c - d
	fmt.Println("c:", c)
	c *= d // Equivalent to c = c * d
	fmt.Println("c:", c)
	c /= d // Equivalent to c = c / d
	fmt.Println("c:", c)
	c %= d // Equivalent to c = c % d
	fmt.Println("c:", c)

	// Bitwise operators
	e := 5
	f := 3
	fmt.Println("AND:", e&f)
	fmt.Println("OR:", e|f)
	fmt.Println("XOR:", e^f)
	fmt.Println("AND NOT:", e&^f)
	fmt.Println("Left Shift:", e<<1)
	fmt.Println("Right Shift:", e>>1)
}
