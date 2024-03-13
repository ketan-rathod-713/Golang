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

func fizz(){

}

func buzz(){

}

func fizzbuzz(){

}

// agar tino mese kuch na aaye then after some timeout print normal number
