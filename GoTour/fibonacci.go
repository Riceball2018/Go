package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	sum := -1
	prevNum := 0
	return func() int {
		if sum <= 0 {
			prevNum = sum
			sum += 1	
		} else {
			//temp := sum
			prevNum, sum = sum, prevNum + sum
			// = temp
		}
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}