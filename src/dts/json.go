package dts

import (
	"cnst"
	"encoding/json"
	"encoding/xml"
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

func encodeJSON(b *cnst.Blog, f *os.File) (err error) {
	enc := json.NewEncoder(f)
	err = enc.Encode(b)
	return
}

func encodeXML(b *cnst.Blog, f *os.File) (err error) {
	enc := xml.NewEncoder(f)
	err = enc.Encode(b)
	return
}

func decodeJSON(b *cnst.Blog, f *os.File) (err error) {
	dec := json.NewDecoder(f)
	err = dec.Decode(b)
	return
}

func decodeXML(b *cnst.Blog, f *os.File) (err error) {
	dec := xml.NewDecoder(f)
	err = dec.Decode(b)
	return
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

	err = encodeJSON(&cnst.DummyBlog, f)
	if err != nil {
		logger.Fatalf("%s", err.Error())
		os.Exit(1)
	}

	// Set the pointer to start
	f.Seek(0, 0)

	var myBlog cnst.Blog
	err = decodeJSON(&myBlog, f)

	if err != nil {
		logger.Fatalf("%s", err.Error())
		os.Exit(1)
	}

	fmt.Printf(" Read: %s\n", myBlog.String())

	// Clear the contents
	f.Truncate(0)
	f.Seek(0, 0)

	err = encodeXML(&cnst.DummyBlog, f)
	if err != nil {
		logger.Fatalf("%s", err.Error())
		os.Exit(1)
	}

	// Set the pointer to start
	f.Seek(0, 0)

	var xmlBlog cnst.Blog
	err = decodeXML(&xmlBlog, f)

	if err != nil {
		logger.Fatalf("%s", err.Error())
		os.Exit(1)
	}

	fmt.Printf(" Read: %s\n", xmlBlog.String())
	fmt.Println(" --- End Program -- ")
}
