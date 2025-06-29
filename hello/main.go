package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// a slice of names
	names := []string{"Maximus", "Tiberius", "Aurelius", "Alessius"}

	// request greeting messages for the names
	messages, err := greetings.Hellos(names)
	// if an error was returned, print it to the console
	// and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned message
	// to the console
	fmt.Println(messages)
}
