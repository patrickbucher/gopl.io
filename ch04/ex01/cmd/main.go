package main

import (
	"crypto/sha256"
	"fmt"

	"gopl.io/ch04/ex01"
)

func main() {
	a, b := sha256.Sum256([]byte("hello")), sha256.Sum256([]byte("Hello"))
	fmt.Println(ex01.BitDiff(a[:], b[:]))
}
