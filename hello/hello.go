package main

import (
	"fmt"

	"github.com/mytkdals93/gotutorial/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Sangmin")
	fmt.Println(message)
}
