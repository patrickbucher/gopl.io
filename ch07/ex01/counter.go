package counter

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	var n int
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for {
		ok := scanner.Scan()
		err := scanner.Err()
		if err != nil {
			return n, fmt.Errorf("error counting words: %v", err)
		}
		if ok {
			*c++
			n += len(scanner.Bytes())
		} else {
			break
		}
	}
	return n, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	var n int
	scanner := bufio.NewScanner(bytes.NewReader(p))
	for {
		ok := scanner.Scan()
		err := scanner.Err()
		if err != nil {
			return n, fmt.Errorf("error counting lines: %v", err)
		}
		if ok {
			*c++
			n += len(scanner.Bytes())
		} else {
			break
		}
	}
	return n, nil
}
