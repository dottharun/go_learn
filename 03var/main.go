package main

import "fmt"

func main() {
	//STRINGS

	//a string variable in golang - with explicit typing
	var name1 string = "free"
	var name2 = "fastyping" //exactly same as before but type is inferred
	var name3 string

	name1 = "gtr" //thus reassignable

	fmt.Println(name1, name2, name3, "//")

	name1 = "cool"
	name3 = "bruh"

	fmt.Println(name1, name2, name3, "//")

	// Shorthand notation to declare "variables"
	// EXACTLY SAME AS BEFORE USING `var` keyword except can only be used in a fn
	// but var declaration can be used anywhere we want
	name4 := "yoshi"

	fmt.Println(name4, "//")

	//INT
	var num1 int = 56
	var num2 = 67
	num3 := 890

	fmt.Println(num1, num2, num3)

	//bits and memory
	// var num4 int8 = 9999 //will give error as outside the range
	var num4 int8 = 127 //no err

	fmt.Println(num4)

	var score1 float32 = 25.98
	var score2 float64 = 1965385877.5
	var score3 = 1.5 // inferred as float64

	fmt.Print(score1, score2, score3)
}
