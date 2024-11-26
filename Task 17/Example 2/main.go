package main

import (
	"example/flags"
	"flag"
	"fmt"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func main() {
	userColor := flag.Bool("color", false, "display colorised output")
	flag.Parse()

	if *userColor {
		colorize(ColorBlue, "Hello Digital Ocean")
		return
	}

	fmt.Println("Hello Digital Ocean")

	// flags.FlagsInGo()

}

// why above function not working
