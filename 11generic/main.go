package main

import "fmt"

/*
TYPE PARAMETERS
-  The type parameters of a function appear between brackets,
    before the function's arguments.
*/

// Index returns the index of givenElem in the givenSlice, or -1 if not found.
func Index[T comparable](
	givenSlice []T, givenElem T) int {

	for i, curr := range givenSlice {
		// curr and givenElem are type T, which has the comparable constraint,
		// so we can use == here.
		if curr == givenElem {
			return i
		}
	}
	return -1
}

func TypeParamExample() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}

/* GENERIC TYPES */

type List[T any] struct {
	next *List[T]
	val  T
}

// TODO: add functionality later

func main() {
	TypeParamExample()
}
