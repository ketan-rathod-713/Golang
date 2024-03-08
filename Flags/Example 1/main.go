package main

import (
	"flag"
	"fmt"
)

type customInt int

func (c *customInt) String() string {
	return fmt.Sprintf("%d", *c)
}

func (c *customInt) Set(value string) error {
	// Parse the string value and set the custom flag
	// For simplicity, we'll assume it's an integer
	var intValue int
	_, err := fmt.Sscanf(value, "%d", &intValue)
	if err != nil {
		return err
	}
	*c = customInt(intValue)
	return nil
}

func main() {
	// var nFlag = flag.Int("n", 1, "get count of n")
	// fmt.Println(*nFlag)

	// fmt.Println(*nFlag)

	// fs := flag.NewFlagSet("boolfunc", flag.ContinueOnError)

	// it sets given flag and gives the value of the flag set.
	flag.BoolFunc("some", "usage ", func(s string) error {
		fmt.Println("Message", s)
		return nil
	})

	var n customInt
	flag.Var(&n, "flag", "just a usage")

	flag.Parse()

	fmt.Println(n)
	// number of arguments left after parsing
	fmt.Println(flag.NArg())

	// number of flags set
	fmt.Println(flag.NFlag())
}

// After parsing, the arguments following the flags are available as the slice flag.Args or individually as flag.Arg(i). The arguments are indexed from 0 through flag.NArg-1.
// Flag parsing stops after first non flag argument or terminator --

// example go run main.go -some=fal -flag=a204abc

// Information

/*
Yes, that's correct. When you define a custom type to be used with the flag package, it needs to implement the flag.Value interface. This interface requires implementing two methods: String() and Set(string) error.

String() Method: This method returns a string representation of the value. It's used by the flag package when printing the default value of the flag.

Set(string) error Method: This method parses the string argument and sets the value of the flag accordingly. It's called by the flag package when parsing command-line arguments.

Implementing these methods correctly allows your custom type to be used seamlessly with the flag package, making it behave like the built-in types (int, string, etc.) when parsing flags from the command line.

So, whenever you define a custom type to be used with flags, ensure it satisfies the flag.Value interface by implementing these methods.
*/
