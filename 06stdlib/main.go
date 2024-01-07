package main

import (
	"fmt"
	"strings"
)

func main() {
	greet := "hello friends"

	fmt.Println(strings.Contains(greet, "hello"))
	fmt.Println(strings.ReplaceAll(greet, "hello", "hi"))
	fmt.Println(strings.ToUpper(greet))
}
