// ALTERNATE 1

// package main

// import "fmt"

// // jo bhi output hoga vo iss channel me dal dega
// func fizz(n int, outputChannel chan string) chan string { // handle multiple of 3

// 	if n%3 == 0 {
// 		outputChannel <- "Fizz"
// 	} else {
// 		outputChannel <- ""
// 	}

// 	return outputChannel
// }

// func buzz(n int, outputChannel chan string) chan string { // handle multiple of 5

// 	if n%5 == 0 {
// 		outputChannel <- "Buzz"
// 	} else {
// 		outputChannel <- ""
// 	}

// 	return outputChannel
// }

// func fizzBuzz(n int, outputChannel chan string) chan string { // multiple of both 5 and 3 means 15
// 	// run as a fan in function
// 	fizzChannel := make(chan string)
// 	buzzChannel := make(chan string)

// 	go fizz(n, fizzChannel)
// 	go buzz(n, buzzChannel)

// 	// check output of both fizz and buzz
// 	fizzString := <-fizzChannel
// 	buzzString := <-buzzChannel

// 	finalString := fmt.Sprintf("%v %v%v", n, fizzString, buzzString)

// 	if finalString == "" {
// 		outputChannel <- fmt.Sprintf("%v", n)
// 	} else {
// 		outputChannel <- finalString
// 	}
// 	return outputChannel
// }

// // create 3 channels and go from 1 to 30 and push to it.
// func main() {
// 	var c chan string = make(chan string)

// 	for i := 1; i <= 30; i++ {
// 		go fizzBuzz(i, c)
// 	}

// 	for i := 0; i < 30; i++ {
// 		fmt.Println(<-c)
// 	}
// }

// ALTERNATE 2
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type Data struct {
// 	n      int
// 	Output chan string
// }

// var (
// 	multiple3 chan string
// )

// func fizz(data Data) {
// 	defer wg.Done()

// 	if data.n%3 == 0 {
// 		data.Output <- "Fizz"
// 	}
// }

// func buzz(data Data) {
// 	defer wg.Done()

// 	if data.n%5 == 0 {
// 		data.Output <- "Buzz"
// 	}
// }

// func fizzBuzz(data Data) {
// 	if(data.n % 3 == 0 && data.n % 5 == 0){
// 		data.Output <- "FizzBuzz"
// 	}
// }

// var wg sync.WaitGroup

// func main() {
// 	for i := 1; i <= 30; i++ {
// 		data := Data{n: i, Output: make(chan string)}
// 		wg.Add(2)
// 		go fizz(data)
// 		go buzz(data)

// 		wg.Wait()
// 		fmt.Println(<-data.Output)
// 	}
// }

// FINAL OUTPUT
package main

import (
	"fmt"
	"time"
)

var (
	m3   chan int  = make(chan int)
	m5   chan int  = make(chan int, 0)
	m15  chan int  = make(chan int)
	next chan bool = make(chan bool)
)

func fizz() {
	for {
		// wait for multiple of 3
		data := <-m3
		fmt.Println("Fizz for ", data)

		// Now next goroutine can start processing
		next <- true
	}
}

func buzz() {
	for {
		// wait for multiple of 5
		data := <-m5
		fmt.Println("Buzz for ", data)

		next <- true
	}
}

func fizzbuzz() {
	for {
		// wait for multiple of 15
		data := <-m15
		fmt.Println("FizzBuzz for ", data)

		next <- true
	}
}

func main() {
	start := time.Now()
	go fizz()
	go buzz()
	go fizzbuzz()

	for i := 1; i <= 30; i++ {
		if i%15 == 0 {
			m15 <- i
			<-next // Don't Procceed Further Until FizzBuzz goruotine Performs it's execution.

		} else if i%5 == 0 {
			m5 <- i
			<-next

		} else if i%3 == 0 {
			m3 <- i
			<-next

		} else {
			fmt.Println(i)
		}

		// don't procceed to next execution before permiss
	}

	elapsed := time.Now().Sub(start)
	fmt.Println("Time Taken", elapsed)
}
