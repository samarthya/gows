package cnst

import (
	"fmt"
)

// Blog blog structure
type Blog struct {
	Title, Author, Category string
	Length                  int
}

func (b *Blog) String() string {
	return fmt.Sprintf(" Blog: %s @ %s\n Category: %s\n Length: %d", b.Title, b.Author, b.Category, b.Length)
}
