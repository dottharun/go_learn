package main

import (
	"fmt"
	"math"
	"strings"
)

func ptr() {
	i, j := 42, 2701

	//pointers type variables hold memory address of a variable

	// `p` is a pointer which holds the memory address of i
	// thus p points to i
	//same as
	// var p *int = &i
	p := &i

	// read i through the pointer p using the dereferencing operator
	fmt.Println(*p)

	// Access and set i through its pointer p
	// Using deferencing operator `*p` is the same as `i`
	*p = 999
	fmt.Println(i) // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	/*
	   - dereferencing aka indirecting
	   - no pointer arithmetic unlike C/C++

	   Note: In golang: pointer types are written as *<type>, *int, *string etc
	   But in C/C++: pointer types are written as <type>*, int*, char*
	*/
}

// structExample: It's a collection of fields
type Vertex struct {
	X int
	Y int
}

func structExample() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 444
	fmt.Println(v)

	//** Pointer to structs doesnt require deference operator to access the
	//      fields of the structs, just `.` is enough
	//      No need this (*p).X just this p.X is enough
	p := &v
	p.X = 777
	fmt.Println(v)

	//Struct Literals - Name: syntax
	//
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p2 = &Vertex{1, 2} // has type *Vertex ---ptr to an anonymous Vertex struct
	)

	fmt.Println(v1, v2, v3, p2)
}

func ArrExample() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func SliceExample() {
	//this is an array
	primes := [6]int{2, 3, 5, 7, 11, 13}

	//this is a slice (using the above array ---references the same array no new arr is created)
	var mySlice []int = primes[1:4]
	fmt.Println(mySlice)

	//Demo to show
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names) //names is changes by its refered slice b

	//SLICE LITERALS - creates an array and makes an slice referencing it
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	fmt.Println("Hello, slicing")

	//SLICE CAPACITY and LENGTH
	//length is the number of elements slice contains
	//capacity is the number of elements in the underlying array, counting from the first element in the slice
	s2 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s2, len(s2), cap(s2)) //len: 6, cap: 6

	// Slice the slice to give it zero length.
	s2 = s2[:0]
	fmt.Println(s2, len(s2), cap(s2)) //len:0, cap:6

	// Extend its length.
	s2 = s2[:4]
	fmt.Println(s2, len(s2), cap(s2)) //len:4, cap:6

	// Drop its first two values. --now capacity also changes since count is started from initial idx
	s2 = s2[2:]
	fmt.Println(s2, len(s2), cap(s2)) //len:2, cap:4

	//NIL slice
	//A nil slice has a length and capacity of 0 and has no underlying array
	var s3 []int
	fmt.Println(s3, len(s3), cap(s3))
	if s3 == nil {
		fmt.Println("Its a nil! slice")
	}

	//Slice with make
	//second field is for length of slice and second field is for capacity - len of underlying array
	a2 := make([]int, 5)
	printSlice("a2", a2)

	b2 := make([]int, 0, 5)
	printSlice("b2", b2)

	c2 := b2[:2]
	printSlice("c2", c2)

	d2 := c2[2:5]
	printSlice("d2", d2)
}

func printMap(mp map[string]int) {
	for k, v := range mp {
		fmt.Println(k, v)
	}
}

func wordCount(s string) map[string]int {
	words := strings.Split(s, " ")

	mp := make(map[string]int)
	for _, w := range words {
		mp[w]++
	}

	return mp
}

func MapExample() {
	wc := wordCount("where are you now where now man where bruh pls")
	printMap(wc)
}

// function value example
// funcs are just values too
func Points3And4Compute(myFunc func(float64, float64) float64) float64 {
	res := myFunc(3, 4)
	return res
}

func FunctionValueExample() {
	//hypot is a variable of type function -- a function(inside another function)
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	z := Points3And4Compute(hypot)
	fmt.Println(z)
}

func Adder() func(int) int {
	//Here sum acts as the member variable of the OOP class Adder
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

// makes a func that gives the fib series number each time its called
func fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		res := a
		a, b = b, a+b
		return res
	}
}

// A closure is a function value that references variables from outside its body
// ---closure functions are just member functions of some class in OOP - here Adder is the class
// and the returned function is the member function
func ClosureExample() {
	//here pos and neg are separate instances of the class Adder thus separate sum member for each
	pos, neg := Adder(), Adder()

	for i := 0; i < 10; i++ {
		fmt.Println("Pos:", pos(i), "Neg:", neg(-2*i))
	}

	// Exercise Fibonacci closure
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func main() {
	// ptr()
	// structExample()
	// ArrExample()
	// SliceExample()
	// MapExample()
	// FunctionValueExample()
	ClosureExample()
}
