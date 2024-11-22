// // Original Program: How can we reduce the execution time of this program using concurrency?
// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"
// )

// func task(i int) {
// 	// Task function simulates some work by printing a message and then sleeping for 100 milliseconds
// 	fmt.Println("Task", i)
// 	time.Sleep(100 * time.Millisecond)
// }

// func main() {
// 	start := time.Now()

// 	// Loop through 30 tasks sequentially
// 	for i := 1; i <= 30; i++ {
// 		task(i)
// 	}

// 	elapsed := time.Since(start)
// 	log.Printf("Time taken %s", elapsed)
// }

// OUTPUT

package main

import (
	"fmt"
	"log"
	"time"
)

var (
	c chan string = make(chan string)
)

func task(i int) {
	// Task function simulates some work by printing a message and then sleeping for 100 milliseconds
	// fmt.Println("Task", i)
	c <- fmt.Sprintf("Task %d", i)
	time.Sleep(100 * time.Millisecond)
}

func main() {
	start := time.Now()

	// Loop through 30 tasks sequentially
	for i := 1; i <= 30; i++ {
		go task(i)
	}

	// now take output using channel it will be blocking by default
	for i := 1; i <= 30; i++ {
		fmt.Println(<-c)
	}

	close(c)

	elapsed := time.Since(start)
	log.Printf("Time taken %s", elapsed)
}
