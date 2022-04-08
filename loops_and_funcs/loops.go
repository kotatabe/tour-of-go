package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	fmt.Println("== z ==")
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Println(i+1, ":",z)
	}
	return z
}

func main() {
	x := 999.0
	fmt.Println("My Sqrt   :", Sqrt(x))
	fmt.Println("math.Sqrt :", math.Sqrt(x))
}
