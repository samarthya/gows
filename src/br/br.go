package br

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var fileName string

func init() {
	flag.StringVar(&fileName, "f", "", "File to read")
	flag.Parse()
	fmt.Println(" Init: fileName: " + fileName)
}

const (
	// InvalidArgument for the program
	InvalidArgument = "Please provide a valid file\n"
)

// Cmd Runs the command
func Cmd() {
	// https://golang.org/pkg/flag/#NArg
	// NArg is the number of arguments remaining after flags have been processed.
	fmt.Fprintf(os.Stdout, "DBG: %d %s\n", flag.NArg(), fileName)

	if flag.NFlag() == 0 {
		fmt.Fprintf(os.Stderr, "%s", InvalidArgument)
		flag.PrintDefaults()
		os.Exit(1)
	}

	//OpenAndDump(flag.Arg(i))
	OpenAndDump(fileName)

}

//dump dumps the content from a reader
func dump(fr *bufio.Reader) {
	fmt.Println("\n -------- STARTOFFILE -------")
	for {
		b, err := fr.ReadByte()
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, " %X:%c", b, b)
	}
	fmt.Println("\n ------ ENDOFFILE ------")
}

// OpenAndDump a file and dumps the content
func OpenAndDump(f string) {
	fh, err := os.Open(f)
	if err == nil {
		dump(bufio.NewReader(fh))
		os.Exit(0)
	}

	defer fh.Close()
	fmt.Fprintf(os.Stderr, "ERR: (%s) %s", f, err.Error())
	return
}
