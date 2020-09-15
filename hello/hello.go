package main

import (
	"fmt"
	"log"

	"github.com/mytkdals93/gotutorial/greetings"
)

func main() {

	// Get a greeting message and print it.
	message, _ := greetings.Hello("Sangmin")

	fmt.Println(message)
	//executeGrettingWithError()
	greetingMultipleNames()
}
func executeGrettingWithError() {
	// Set properties of the predefined Logge, incliding
	// the log entry prefix and a flag to disabe printing
	// the time, source file, and line number.
	log.SetPrefix("greeting.Hello: ")
	log.SetFlags(0)

	// Request a greeting message with empty string.
	message, err := greetings.Hello("")
	//If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returend, print the returned message
	// to the console
	log.Println(message)
}
func greetingMultipleNames() {
	log.SetPrefix("greeting.Hellos(): ")
	log.SetFlags(0)
	names := []string{"Gladys", "Samantha", "Darrin"}
	g, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(g)
}
