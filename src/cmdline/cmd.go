package cmd

import (
	"flag"
	"fmt"
	"os"
)

var nTimes int

func init() {
	flag.IntVar(&nTimes, "n", 1, "number of times the message needs to be displayed")
}
// A banner command that will display a message an x number of times
// message - string
// times - int
func main() {

	var msg = flag.String("m", "", "the message that needs to be printed")
	var dbg = flag.Bool("d", false, "to debug or not")

	// Must be called after all flags are defined and before flags are accessed within the program
	flag.Parse()

	if *dbg {
		fmt.Printf("DBG: Number of Arguments: %d", len(os.Args))
	}

	if nTimes <= 0 && *dbg {
		fmt.Println("value should be always more than equal to 1")
		os.Exit(1)
	} else {

		if len(*msg) == 0 {
			flag.PrintDefaults()
			os.Exit(2)
		}

		fmt.Println("\n--- Command line Program ---")
		for i := nTimes; i > 0; i-- {
			fmt.Fprintf(os.Stdout, "\n >> %s\n", *msg)
		}
	}
}
