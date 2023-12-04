package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting string for the name string
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// create a msg using random format
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	var formats = []string{
		"Hi, %v, Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
