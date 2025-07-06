package main

import (
	"fmt"
	"rsc.io/quote"
	"time"
)

// greeting returns a pleasant, semi-useful greeting.
func greeting() string {
	return "Hello world, the time is: " + time.Now().String()
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(greeting())
}
