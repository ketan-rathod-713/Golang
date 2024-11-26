package flags

import (
	"flag"
	"log"
)

func FlagsInGo() {
	var count int
	flag.IntVar(&count, "n", 5, "number of lines to read from the file")
	flag.Parse()

	log.Println(count)
}
