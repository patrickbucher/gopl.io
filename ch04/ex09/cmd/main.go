package main

import (
	"fmt"
	"os"

	"gopl.io/ch04/ex09"
)

func main() {
	wordfreq := ex09.Wordfreq(os.Stdin)
	for k, v := range wordfreq {
		fmt.Printf("%9d\t%s\n", v, k)
	}
}
