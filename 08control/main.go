package main

import "fmt"

func givesomething() string {
	defer fmt.Println("oops")
	return "right"
}

func def1() {
	//defer statement means
	//the function is not executed until the surrounding function finishes
	//returning
	//while if the deferred function calls something, they are executed at immediately/sync
	//only the deferred function is delayed
	//---thus very different and simpler than js - where everythings is delayed for async calls
	defer fmt.Println(givesomething())

	//Here first givesomething is executed and right is given to the deferred println in line17
	//and immediatedly after the return all deffered funcs of givesomething is executed thus oops is printed first
	//and then main func hello is printed and only then the returned string from givesomething is printed
	//which is right

	fmt.Println("hello")
}

func def2() {
	// stacking defers defers are pushed onto a stack
	// thus defers are finally called in last in first out manner
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	//countdown is ran from 9, 8, 7, ..., 1, 0
}

func main() {
	// def1()
	def2()
}
