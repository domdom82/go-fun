package main

import (
	"fmt"
	"os"
	"strings"
)

// This is my first ever go program. Excited to learn a new language!
func main() {
	name := "World"

	// If a name was passed, we greet that name instead.
	if len(os.Args) > 0 {
		name = strings.Join(os.Args[1:], " ")
	}

	fmt.Printf("Hello %s!\n", name)
}
