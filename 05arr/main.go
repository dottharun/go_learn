package main

import (
	"fmt"
)

func main() {
	// arrays - are statically allocated => no dynamic length/allocation => fixed length
	var ages [3]int = [3]int{20, 25, 34}
	var ages2 = []int{20, 25, 34}

	names := [4]string{"yoshi", "mario", "peach", "bowser"}

	fmt.Println("nice", ages, ages2, len(ages))
	fmt.Println(names)

	// slices - length is dynamic
	// Here an fixed size array is made first and then a slice is created for that array
	//mostly used instead of arrays
	var scores = []int{100, 50, 60}
	scores[2] = 99

	scores = append(scores, ages[:]...)

	fmt.Println(scores)

	//slice expressions - with a given range
	//converts a given array/slice to a ranged slice
	//beginning and ending is optional
	range1 := names[1:3]
	range2 := names[:3]
	range3 := names[1:]

	fmt.Println(range1, range2, range3)

	//since they are just slices we can do some append operations now
	range2 = append(range2, "koopa")

	fmt.Println(range2, range3)
}
