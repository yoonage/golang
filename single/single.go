package single

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name, family string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf("Hi, %v %v. Welcome!", name, family)
	return message, nil
}
