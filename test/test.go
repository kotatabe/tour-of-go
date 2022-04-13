package main

import (
	"fmt"
	"time"
)

func putstr(ch chan int) {
	fmt.Println(<-ch)
	time.Sleep(time.Second * 3)
	fmt.Println(<-ch)
	fmt.Println("a")
}

func main() {
	ch := make(chan int, 1)
	
	go func(){
		ch <- 100
		ch <- 200
		fmt.Println("send")
	}()
	putstr(ch)
}