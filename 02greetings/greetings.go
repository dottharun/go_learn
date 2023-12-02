package greetings

import (
	"errors"
	"fmt"
)

// returns a greeting string for the name string
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	//Returns a greeting that embeds a name in a message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message, nil
}
