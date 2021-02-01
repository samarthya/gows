package cry

import (
	"crypto"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
)

// Cmd main entry point
func Cmd() {
	fmt.Printf(" Is BLAKE2b_256 available: %t\n", crypto.BLAKE2b_256.Available())
	fmt.Printf(" Is MD5 available: %t\n", crypto.MD5.Available())
	fmt.Printf(" Is SHA1 available: %t\n", crypto.SHA1.Available())
	fmt.Printf(" Is SHA256 available: %t\n", crypto.SHA256.Available())
	var hasher hash.Hash
	// Creates a new MD5 hasher
	hasher = md5.New()

	// Version 1
	io.WriteString(hasher, "And Leon's getting laaarger!")
	fmt.Printf("%x\n", hasher.Sum(nil))
	fmt.Printf("Size: %d\n", hasher.Size())

	// Version 2
	var data = []byte("Data for which the md5 hash will be computed")
	fmt.Printf("%x\n", hasher.Sum(data))
	fmt.Printf("Size: %d\n", hasher.Size())

	sha1 := sha1.New()
	fmt.Printf("%x\n", sha1.Sum(nil))
	fmt.Printf("Size: %d\n", sha1.Size())

	sha256 := sha256.New()
	fmt.Printf("%x\n", sha256.Sum(nil))
	fmt.Printf("Size: %d\n", sha256.Size())
}
