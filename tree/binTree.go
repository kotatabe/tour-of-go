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
	// closeするのか　→しなくて良い
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// treeの最大値がわからない場合は？
	// 	「size_maxを返す関数を作る」しか思いつかなかったです。
	x, y := 0, 0
	for i := 0; i < 10; i++ {
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
	arg := 2

	fmt.Println("==== Test Walk(", arg, ") ====")
	go Walk(tree.New(arg), ch)
	for i := 0; i < 10; i++ {
		fmt.Print(<-ch, " ")
	}

	arg1 := 2
	arg2 := 3
	fmt.Println("\n", "==== Test Same(", arg1, ",", arg2, ") ====")	
	if b := Same(tree.New(arg1), tree.New(arg2)); b {
		fmt.Println("Same!")
	} else {
		fmt.Println("Not same..")
	}
}
