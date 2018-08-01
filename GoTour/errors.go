package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)	
	}
	
	z := x / 2
	z = 1.0
	i := 0
	for math.Abs(z*z - x) > 0.001 {
		z -= (z*z - x) / (2*z)
		i += 1
		//fmt.Println("z is", z)
	}
	//fmt.Println(i, "iteration(s)")
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}