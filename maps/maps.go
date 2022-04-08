package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	splited_ar := strings.Fields(s)
	word_map := make(map[string]int)

	for _, string := range splited_ar {
		word_map[string] += 1
	}
	fmt.Println(word_map)
	return word_map
}

func main() {
	wc.Test(WordCount)
}
