package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	z = 1.0
	i := 0
	for math.Abs(z*z - x) > 0.001 {
		z -= (z*z - x) / (2*z)
		i += 1
		fmt.Println("z is", z)
	}
	fmt.Println(i, "iteration(s)")
	return z
}

func main() {
	num := 6.0
	mySqrt := Sqrt(num)
	fmt.Println("Closeness to library version", mySqrt / math.Sqrt(num))
}
