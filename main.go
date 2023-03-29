package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go printOne(c)
	go printTwo(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func printOne(c chan int) {
	time.Sleep(1 * time.Second)
	c <- 1
}

func printTwo(c chan int) {
	c <- 2
}
