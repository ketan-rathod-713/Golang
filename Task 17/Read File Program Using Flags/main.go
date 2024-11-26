package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	var n int
	var m int

	//  This function allows us to pass our own pointer to a variable in contrast to the flag functions that do not have the Var suffix
	flag.IntVar(&n, "n", 5, "Number of lines to read from file")
	flag.IntVar(&m, "m", 5, "Number of lines to read from file")
	flag.Parse() // after all the flags are defined and before any flag is accessed

	fmt.Println("Value of n and m is", n, m)
	fmt.Println(flag.Arg(0))

	// take n in input and a file name
	filename := flag.Arg(0)

	var in io.Reader
	if filename == "" {
		flag.Usage()
		return
	} else {
		// read 5 lines
		file, err := os.Open(filename)
		if err != nil {
			log.Println(err)
			return
		}

		in = file
	}

	buf := bufio.NewScanner(in)

	for i := 0; i < n; i++ {
		if !buf.Scan() {
			break
		}

		fmt.Println(buf.Text())
	}

	if err := buf.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading: err:", err)
	}
}

// flag.Arg function to access the first positional argument after all flags.
// if not set by user then it will be empty string
