package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// get the greeting string and print it
	message := greetings.Hello("colaman")
	fmt.Println(message)
}
