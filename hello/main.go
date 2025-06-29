package main

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	// get a greeting message and print it
	message := greetings.Hello("Amadeus")
	fmt.Println(message)
}
