package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

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
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	x, y := 0, 0
	for i := 0; i < 10; i++{
		x, y = <-ch1, <-ch2
		fmt.Println(x, y)
		if x != y {
			return false	
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	tree_num := 2

	fmt.Println("==== Test Walk(", tree_num, ") ====")
	go Walk(tree.New(tree_num), ch)
	for i := 0; i < 10; i++ {
		fmt.Print(<-ch, " ")
	}

	fmt.Println("\n", "==== Test Same(", tree_num1, ",", tree_num2, ") ====")	
	tree_num1 := 2
	tree_num2 := 3
	if b := Same(tree.New(tree_num1), tree.New(tree_num2)); b {
		fmt.Println("Same!")
	} else {
		fmt.Println("Not same..")
	}
}
