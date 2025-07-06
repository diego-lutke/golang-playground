// code source: https://withcodeexample.com/golang-tutorial-for-beginners/
package main

import "fmt"

func main() {
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

	fmt.Printf("Zero int: %d\n", zeroInt)
	fmt.Printf("Zero float: %f\n", zeroFloat)
	fmt.Printf("Zero string: %s\n", zeroString)
	fmt.Printf("Zero bool: %t\n", zeroBool)
}
