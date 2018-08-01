package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	
	ch <- t.Value
	
	if t.Right != nil {
		Walk(t.Right, ch)	
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	isSame := true
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	// Check 10 values
	for i := 0; i < 10; i++ {
		v1, v2 := <-ch1, <-ch2
		fmt.Println(v1, v2)
		if v1 != v2 {
			isSame = false	
		}
	}
	return isSame	
}

func main() {
	
	fmt.Println(Same(tree.New(2), tree.New(2)))
}