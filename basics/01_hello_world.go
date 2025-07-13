package basics

import (
	"fmt"
	"github.com/enescakir/emoji"
	"rsc.io/quote"
	"time"
)

// greeting returns a pleasant, semi-useful greeting.
func greeting() string {
	return "Hello world, the time is: " + time.Now().String()
}

func HelloWorld() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(greeting())
	fmt.Println(emoji.WavingHand.Tone())
}
