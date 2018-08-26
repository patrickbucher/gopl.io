// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := []struct {
		function func(rune) bool
		category string
		count    int
	}{
		{unicode.IsControl, "control", 0},
		{unicode.IsDigit, "digit", 0},
		{unicode.IsGraphic, "graphic", 0},
		{unicode.IsLetter, "letter", 0},
		{unicode.IsLower, "lower", 0},
		{unicode.IsMark, "mark", 0},
		{unicode.IsNumber, "number", 0},
		{unicode.IsPrint, "print", 0},
		{unicode.IsPunct, "punct", 0},
		{unicode.IsSpace, "space", 0},
		{unicode.IsSymbol, "symbol", 0},
		{unicode.IsTitle, "title", 0},
		{unicode.IsUpper, "upper", 0},
	}
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		for i := range counts {
			if counts[i].function(r) {
				counts[i].count++
			}
		}
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for _, c := range counts {
		fmt.Printf("%s\t%5d\n", c.category, c.count)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%5d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
