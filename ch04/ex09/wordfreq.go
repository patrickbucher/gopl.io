package ex09

import (
	"bufio"
	"io"
)

func Wordfreq(r io.Reader) map[string]int {
	wordfreq := make(map[string]int)
	input := bufio.NewScanner(r)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		wordfreq[input.Text()]++
	}
	return wordfreq
}
