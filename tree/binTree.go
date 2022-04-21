package main

import (
	"fmt"
	"sync"
	"golang.org/x/tour/tree"
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
	// closeするのか　→しなくて良い
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go Walk(t2, ch2)
	// ch1だけでよい
	// treeの最大値がわからない場合は？
	for i := range ch1 {
		x, y := i, <-ch2
		fmt.Println(x, y)
		if x != y {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	tree_n := 2

	fmt.Println("==== Test Walk(", tree_n, ") ====")
	go Walk(tree.New(tree_n), ch)
	for i := 0; i < 10; i++ {
		fmt.Print(<-ch, " ")
	}

	tree_n1 := 2
	tree_n2 := 3
	fmt.Println("\n", "==== Test Same(", tree_n1, ",", tree_n2, ") ====")	
	if b := Same(tree.New(tree_n1), tree.New(tree_n2)); b {
		fmt.Println("Same!")
	} else {
		fmt.Println("Not same..")
	}
}
