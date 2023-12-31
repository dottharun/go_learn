package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	//setting the std Logger:
	//prefix for every log call
	//and using flag 0(---enum like thing to where 0 means no flags) to disable logging time and date of the call
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"naveen", "coolio", "Darren Brown"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
