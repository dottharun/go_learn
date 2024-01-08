package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1000.0

	for i := 1; i <= 10; i++ {
		fmt.Println("guess is ", z)
		z -= (z - (x / z)) / 2
	}

	return z
}

func main() {
	fmt.Println(Sqrt(36))
}
