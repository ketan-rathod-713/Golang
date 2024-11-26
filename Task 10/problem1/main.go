package main

import (
	"fmt"
	"time"
)

var (
	even chan int = make(chan int)
	odd  chan int = make(chan int)
)

// using select
func a() {
	for {
		select {
		case <-even:
			fmt.Println(0)
			odd <- 1
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// without using select
func b() {
	for {
		<-odd
		fmt.Println(1)
		even <- 0
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	go a()
	go b()

	// intial seeding to channel
	odd <- 1

	// main function waits for 1 sec before exiting
	time.Sleep(1 * time.Second)
}
