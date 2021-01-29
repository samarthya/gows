package main

import (
	"br"
	"flag"
	"fmt"
	"os"
)

var fileName string

func init() {
	flag.StringVar(&fileName, "-f", "", "File to read")
}

const (
	// InvalidArgument for the program
	InvalidArgument = "Please provide a valid file"
)

func main() {

	flag.Parse()

	// https://golang.org/pkg/flag/#NArg
	// NArg is the number of arguments remaining after flags have been processed.
	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "%s", InvalidArgument)
		os.Exit(1)
	}

	for i := 0; i < flag.NArg(); i++ {
		br.OpenAndDump(flag.Arg(i))
	}
}
