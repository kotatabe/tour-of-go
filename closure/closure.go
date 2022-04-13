package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

func fibonacci() func(int) int {
	before2, before1 := 0,1
	return func(i int) int {
		if i == 0 || i == 1 {
			return i
		}
		ret_int := before2 + before1
		before2 = before1
		before1 = ret_int
		return ret_int
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
