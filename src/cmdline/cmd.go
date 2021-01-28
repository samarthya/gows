package main

import (
	"flag"
	"fmt"
	"os"
)

// A banner command that will display a message an x number of times
// message - string
// times - int
func main() {
	fmt.Println(" --- Command line Program --- ")
	var times = flag.Int("n", 1, "number of times the message needs to be displayed")
	var msg = flag.String("m", "please look at the help for more information\n", "the message that needs to be printed")
	var dbg = flag.Bool("d", false, "to debug or not")

	flag.Parse()

	if *dbg {
		fmt.Printf("DBG: Number of Arguments: %d", len(os.Args))
	}

	if *times <= 0 && *dbg {
		fmt.Println("value should be always more than equal to 1")
		os.Exit(1)
	} else {
		for i := *times; i > 0; i-- {
			fmt.Fprintf(os.Stdout, "\n >> %s\n", *msg)
		}
	}
}
