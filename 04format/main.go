package main

import "fmt"

func main() {
	fmt.Print("nice")
	fmt.Print("bruh\n") //doesnt print newline automatically

	fmt.Println("welcome")
	name := "sri"
	age := 35

	fmt.Println("My name is", name)
	fmt.Println("My name is", name, "My age is", age) //can have multiple comma separated values

	// formatted strings
	// uses format specifiers `%v` to do formatting variables
	// %q for quotes
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name) //???something weird happening here not sure though
	fmt.Printf("age is of type %T  \n", age)
	fmt.Printf("you scored %f points \n", 225.55)
	fmt.Printf("you scored %0.1f points \n", 225.55)

	//save formatted string
	str := fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println(str)
}
