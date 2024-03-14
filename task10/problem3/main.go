// ALTERNATE 1

package main

import (
	"fmt"
	"time"
)

// output will be returned by output channel
func fizz(n int, outputChannel chan string) { // handle multiple of 3
	time.Sleep(100 * time.Millisecond)

	if n%3 == 0 {
		outputChannel <- "Fizz"
	} else {
		outputChannel <- ""
	}
}

func buzz(n int, outputChannel chan string) { // handle multiple of 5
	time.Sleep(200 * time.Millisecond)

	if n%5 == 0 {
		outputChannel <- "Buzz"
	} else {
		outputChannel <- ""
	}
}

func fizzBuzz(n int, c chan string) { // multiple of both 5 and 3 means 15

	time.Sleep(50 * time.Millisecond)

	// run as a fan in function
	fizzChannel := make(chan string)
	buzzChannel := make(chan string)

	go fizz(n, fizzChannel)
	go buzz(n, buzzChannel)

	// check output of both fizz and buzz
	// Waits for both fizz and buzz channel.
	fizzString := <-fizzChannel
	buzzString := <-buzzChannel

	finalString := fmt.Sprintf("%v%v", fizzString, buzzString)

	if finalString == "" {
		// do nothing
	} else {
		fmt.Printf("%v, %v\n", n, finalString)
	}
	c <- "wow"
}

// create 3 channels and go from 1 to 30 and push to it.
func main() {
	startTime := time.Now()
	var c chan string = make(chan string)

	for i := 1; i <= 30; i++ {
		go fizzBuzz(i, c)
	}

	for i := 0; i < 30; i++ {
		<-c
	}

	elapsedTime := time.Now().Sub(startTime)
	fmt.Println("Time Taken", elapsedTime)
}

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
// package main

// import (
// 	"fmt"
// 	"time"
// )

// var (
// 	m3   chan int  = make(chan int)
// 	m5   chan int  = make(chan int, 0)
// 	m15  chan int  = make(chan int)
// 	next chan bool = make(chan bool)
// )

// func fizz() {
// 	for {
// 		// wait for multiple of 3
// 		data := <-m3
// 		fmt.Println("Fizz for ", data)

// 		// Now next goroutine can start processing
// 		next <- true
// 	}
// }

// func buzz() {
// 	for {
// 		// wait for multiple of 5
// 		data := <-m5
// 		fmt.Println("Buzz for ", data)

// 		next <- true
// 	}
// }

// func fizzbuzz() {
// 	for {
// 		// wait for multiple of 15
// 		data := <-m15
// 		fmt.Println("FizzBuzz for ", data)

// 		next <- true
// 	}
// }

// func main() {
// 	start := time.Now()
// 	go fizz()
// 	go buzz()
// 	go fizzbuzz()

// 	for i := 1; i <= 30; i++ {
// 		if i%15 == 0 {
// 			m15 <- i
// 			<-next // Don't Procceed Further Until FizzBuzz goruotine Performs it's execution.

// 		} else if i%5 == 0 {
// 			m5 <- i
// 			<-next

// 		} else if i%3 == 0 {
// 			m3 <- i
// 			<-next

// 		} else {
// 			fmt.Println(i)
// 		}

// 		// don't procceed to next execution before permiss
// 	}

// 	elapsed := time.Now().Sub(start)
// 	fmt.Println("Time Taken", elapsed)
// }

// NOTE : put core logic inside function and channels should be local and not global one
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type data struct {
// 	n    int
// 	text string
// }

// func fizz(n int, checkFizzBuz chan data, wg sync.WaitGroup) {
// 	defer wg.Done()
// 	if n%3 == 0 {
// 		// Now send this data for fizzbuzz check
// 		checkFizzBuz <- data{n: n, text: "Fizz"}
// 	}
// }

// func buzz(n int, checkFizzBuz chan data, wg sync.WaitGroup) {
// 	defer wg.Done()
// 	if n%5 == 0 {
// 		// Now send this data for fizzbuzz check
// 		checkFizzBuz <- data{n: n, text: "Buzz"}
// 	}
// }

// func fizzBuzz(checkFizzBuz chan data, wg sync.WaitGroup) {
// 	// Get Data From Fizzbuz check channel and process it accordingly.
// 	defer wg.Done()

// 	for {
// 		wg.Add(1)
// 		data, ok := <-checkFizzBuz
// 		if !ok {
// 			break
// 		} else {
// 			if data.n%3 == 0 && data.n%5 == 0 {
// 				fmt.Printf("%v,%v\n", data.n, "FizzBuzz")
// 			} else {
// 				fmt.Printf("%v,%v\n", data.n, data.text)
// 			}
// 		}
// 	}
// }

// func main() {
// 	var checkFizzBuz chan data = make(chan data)
// 	var wg sync.WaitGroup

// 	go fizzBuzz(checkFizzBuz, wg)

// 	for i := 1; i <= 30; i++ {
// 		wg.Add(2)
// 		go fizz(i, checkFizzBuz, wg)
// 		go buzz(i, checkFizzBuz, wg)
// 	}

// 	wg.Wait()
// }
