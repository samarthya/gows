package enc

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"os"
)

// Blog blog structure
type Blog struct {
	Title, Author, Category string
	Length                  int
}

func (b *Blog) String() string {
	return fmt.Sprintf(" Blog: %s @ %s\n Category: %s\n Length: %d", b.Title, b.Author, b.Category, b.Length)
}

var fileName string
var logger = log.New(os.Stdout, " ENC: ", log.LUTC)

func init() {
	flag.StringVar(&fileName, "f", "", "Output/Input file")
	flag.Parse()
}

const (
	//ErrorMessage For the errors
	ErrorMessage = "need a file to write or read the contents."
)

var defBlog = Blog{Title: "My first blog", Author: "Saurabh Sharma", Category: "technical", Length: 200}

// Cmd main entry method
func Cmd() {
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		logger.Fatalf("%s", ErrorMessage)
	}

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0666)

	defer f.Close()

	if err != nil {
		logger.Fatalf("%s", err.Error())
		os.Exit(1)
	}

	enc := gob.NewEncoder(f)
	err = enc.Encode(&defBlog)
	if err != nil {
		logger.Fatalf("%s", err.Error())
		os.Exit(1)
	}

	dec := gob.NewDecoder(f)
	var blogR Blog

	f.Seek(0,0)
	err = dec.Decode(&blogR)
	if err != nil {
		logger.Fatalf("%s", err.Error())
		os.Exit(1)
	}

	fmt.Printf(" Read: %s\n", blogR.String())
	fmt.Println(" --- End Program -- ")
}
