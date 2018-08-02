// Dup2 prints the count and text of lines that appear more than once in the
// input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Occurence struct {
	files map[string]bool
	count int
}

func main() {
	counts := make(map[string]*Occurence)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, oc := range counts {
		if oc.count > 1 {
			fmt.Printf("%s\t%d\n", line, oc.count)
			for f, _ := range oc.files {
				fmt.Printf("\t%s\n", f)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]*Occurence) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		if _, ok := counts[text]; !ok {
			var oc Occurence
			oc.files = make(map[string]bool)
			oc.count = 0
			counts[text] = &oc
		}
		counts[text].files[f.Name()] = true
		counts[text].count = counts[text].count + 1
	}
	// NOTE: ignoring potential errors from input.Err()
}
