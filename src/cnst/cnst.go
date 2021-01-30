package cnst

import (
	"fmt"
)

// Blog blog structure
type Blog struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"cat"`
	Length   int    `json:"size"`
}

func (b *Blog) String() string {
	return fmt.Sprintf(" Blog: %s @ %s\n Category: %s\n Length: %d", b.Title, b.Author, b.Category, b.Length)
}

// DummyBlog blog value for testing
var DummyBlog = Blog{Title: "My first blog", Author: "Saurabh Sharma", Category: "technical", Length: 200}
