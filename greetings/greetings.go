package greetings

// this package will be used as a module
import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// if a name was received return a greeting that embeds the name in a message
	// the := operator is a shortcut for declaring and initializing a variable in one line
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message, nil
}
