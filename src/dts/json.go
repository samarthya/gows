package dts

import (
	"cnst"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

// JSON marshalling
// Wil use the blog struct defined in the source files.
// Have added the tags

var fileName string
var logger = log.New(os.Stdout, " DTS: ", log.LUTC)

func init() {
	flag.StringVar(&fileName, "f", "", "Output/Input file")
	flag.Parse()
}

// Cmd entry point
func Cmd() {
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		logger.Fatalf("%s", "missing -f")
	}

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0666)

	defer f.Close()

	if err != nil {
		logger.Fatalf("%s", err.Error())
		os.Exit(1)
	}
	enc := json.NewEncoder(f)
	err = enc.Encode(cnst.DummyBlog)
	if err != nil {
		logger.Fatalf("%s", err.Error())
		os.Exit(1)
	}

	// Set the pointer to start
	f.Seek(0, 0)
	dec := json.NewDecoder(f)
	var myBlog cnst.Blog
	err = dec.Decode(&myBlog)

	if err != nil {
		logger.Fatalf("%s", err.Error())
		os.Exit(1)
	}

	fmt.Printf(" Read: %s\n", myBlog.String())
	fmt.Println(" --- End Program -- ")
}
