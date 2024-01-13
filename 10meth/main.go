package main

import (
	// "errors"
	"fmt"
	"image"
	"io"
	"math"
	"strings"
	"time"
)

/*
Go does not have classes but "Methods" can be defined on types

- a method is function with a receiver argument(receivers are by default copied not referenced)

- method is just a function symbol

- builtin types can also be receivers

- a method's receiver type must be defined in the same package
    => in OOP the whole of the class must be defined in the same package

*/

type Vertex struct {
	X, Y float64
}

// Here we added a new method for the type Vertex from outside the type
// ---inside the type in C/C++
func (v Vertex) OriginDist() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func MethodExample() {
	v := Vertex{3, 4}
	fmt.Println(v.OriginDist())
}

/*
# POINTER RECEIVER
we can use pointer receivers too, which causes no copy just pointer(memory address passing)

# CHOOSING VALUE RECEIVER OR POINTER RECEIVER:
- pointer receiver has two uses:
1. allows for modification of the receiver data
2. avoids copying the whole receiver on "each method call"

- value receiver use is, it allows for no modification of the receiver with copy

==> In general its recommended to use any one of both without mixing
*/

// the receiver is passed by pointer
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func pointerReceiverExample() {
	v := Vertex{3, 4}
	v.Scale(10)           //method changes the member variable through ptr receiver
	fmt.Println(v.X, v.Y) //and its reflected

	/* as we can see we dont need to pass the receiver object by
	   reference its implicitely allowed automatically */
}

/*
# INTERFACE

- interface type is defined as a set of method symbols

- A value of interface type can hold any value that implements those methods

- used to make generic traits like set of methods for a bunch of types

- Note: direct access to data variables are not possible - maybe through getters and setters


- HOW TO Imagine an Interface type? note: even though we are assigning a variable to our interface variable
    - type of the interface variable is unchanged
    - Interfaces are not normal variables
    - Interfaces have a interface type and a Concrete type of checking variable
    - godocs suggests to imagine an interface as a tuple of these two types side by side on a block
*/

// type MyFloat and its Abs method is defined
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyFloat) Square() float64 {
	return float64(f * f)
}

// type MyVertex and its Abs method is defined
type MyVertex struct {
	X, Y float64
}

// MyVertex only implements Abs as a Pointer receiver, thus no copy is made when calling this method
func (v *MyVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// interface definition implementing Abs() method
type Abser interface {
	Abs() float64
}

func interfaceExample() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := MyVertex{3, 4}

	//Here interface variable a from `Abser` is a copy of MyFloat variable f
	//  with only methods from interface
	a = f
	fmt.Println(a.Abs(), a)

	//Now our interface implements its methods to the reference to MyVertex
	// --observe carefully that even when when interface is made for pointer receiver we cannot make a copy
	a = &v
	// a = v //gives error since it makes copies

	fmt.Println(a.Abs())

	//interface initialized and assigned in same line
	//Here we can clearly see that interfaces are not normal variables
	var b Abser = &v
	fmt.Println(b)
}

/*
Interface for nil values
*/

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func interfaceForNilVariables() {
	var i I

	var t *T //nil variable
	i = t
	describe(i) //(<nil>, *main.T)
	i.M()

	i = &T{"hello"}
	describe(i) //(&{hello}, *main.T)
	i.M()
}

/* Nil interface values */

func nilInterfaceValues() {
	var i I
	describe(i) //(<nil>, <nil>)
	// i.M()    //runtime error
}

/* empty Interface
- Empty interfaces are used by code that handles values of unknown type
- For example, fmt.Print takes any number of arguments of type interface{}
*/

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func emptyInterface() {
	var i interface{}
	describe2(i)

	i = 42
	describe2(i)

	i = "hello"
	describe2(i)
}

/*
# Type Assertions
- provides access to an interface value's underlying concrete value
- by using a user's type assertion of the underlying value's type
- if the assertion fails, the program will panic
*/

func typeAssertions() {
	var i interface{} = "hello"

	// A type assertion on the interface i that the underlying value is string
	//   if failed will panice
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	//when two values are return will not panic
	//and the bool condition will fail
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	fmt.Println(f)
}

/*
# TYPE SWITCHES
- type switches are a construct that permits several type assertions in series
*/

func doSomething(i interface{}) {
	// x := i.(type) //invalid since use .(type) syntax is only allowed with switch blocks

	//.(type) always comes with switch block to know check interfaces(defined and empty interfaces all types)

	switch v := i.(type) {
	case int:
		//here v has type int
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		//here v has type string
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		//here since we checked for type already , go runtime knows the type of v
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func typeSwitches() {
	doSomething(21)
	doSomething("hello")
	doSomething(true)
}

/*
# Stringer:
- given by fmt package
- A stringer type can describe itself as a string with String() method
- the fmt package and other packages use this interface to print values for any custom types

*/

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

func stringerExample() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

/*
# Error Interface
- go errors are expressed with error type values
- the error type is a built-in interface similar to fmt.Stringer with Error() method

- normally an interface is called by using a method on the interface object
*/

// custom type with When and What variables as fields
// we want this custom type when printing errors for some function
type MyErrorData struct {
	When time.Time
	What string
}

// fmt looks for this method for any error interface print string
func (e *MyErrorData) Error() string {
	errStr := fmt.Sprintf("at %v, %s", e.When, e.What)
	return errStr
}

func runSomething() (int, error) {
	// return 3, nil //samething without any error

	// return 3, errors.New("Oh no something is wrong with you ") //For simple string based errors

	// return 3, fmt.Errorf("Oh no something is wrong with you ") //For formatted string

	return 3, &MyErrorData{time.Now(), "it didn't work"}
	//        ^---not clear why we need to pass MyErrorData as reference?
	// NEED MORE INSIGHT???
	//remember pointer receiver methods are implemented to avoid copies of big types
	//thus here too we need to return a reference else copies will be made for this assumed MyVertex type
	//
	//how when passing MyErrorData it suddenly becomes an error interface?
	//See remember interfaces are not normal variables
	//when we set the return type as an interface, the interface contract is self contained in the function itself
}

func errorInterface() {
	x, err := runSomething()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(x)
}

/* ERROR EXERCISE */

type ErrDataNegativeSqrt float64

func (e ErrDataNegativeSqrt) Error() string {
	return fmt.Sprintf(`
ERROR: negative-square-root not implemented yet, maybe try some other libraries/package?
error for number %v`,
		float64(e))
	//    ^--Here conversion must be done to avoid recursion in case of error
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrDataNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func errorExercise() {
	fmt.Println(Sqrt(56))
	fmt.Println(Sqrt(-222))
}

/*
# Reader Interface
- The example code creates a strings.Reader and consumes its output 8 bytes at a time
*/

func ReaderExample() {
	r := strings.NewReader("Homer is here xxxxxxxxxxxxxxxxxxxxxxx")
	b := make([]uint8, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}

/* IMAGE INTERFACE */

func ImageInterface() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

func main() {
	// MethodExample()
	// pointerReceiverExample()
	// interfaceExample()
	// interfaceForNilVariables()
	// nilInterfaceValues()
	// emptyInterface()
	// typeAssertions()
	// typeSwitches()
	// stringerExample()
	// errorInterface()
	// errorExercise()
	// ReaderExample()
    ImageInterface()
}
