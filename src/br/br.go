package br

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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
