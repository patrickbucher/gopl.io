package main

import (
	"fmt"
	"strings"
)

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	input := []string{
		"a",
		"a.go",
		"a/b/c.go",
		"a/b.c.go",
	}
	for i := range input {
		fmt.Printf("%s => %s\n", input[i], basename(input[i]))
	}
}
