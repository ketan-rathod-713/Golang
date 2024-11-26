// Original Program: How can we reduce the execution time of this program using concurrency?
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func task(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Task function simulates some work by printing a message and then sleeping for 100 milliseconds
	fmt.Println("Task", i)
	time.Sleep(100 * time.Millisecond)
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	// Loop through 30 tasks sequentially
	for i := 1; i <= 30; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Time taken %s", elapsed)
}
