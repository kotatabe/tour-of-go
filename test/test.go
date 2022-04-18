package main

import (
	"fmt"
)

func main() {
	str := "abcd"
	for i := 0; str[i] != nil; i++{
		fmt.Println(str[i])
	}
}