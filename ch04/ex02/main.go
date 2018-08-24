package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var size = flag.Int("s", 256, "The hash length, one of 256 (default), 384 or 512")

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Fprintln(os.Stdout, "missing argument for hashing")
		os.Exit(1)
	}
	arg := []byte(flag.Arg(0))
	var sum []byte
	switch *size {
	case 256:
		s := sha256.Sum256(arg)
		sum = s[:]
	case 384:
		s := sha512.Sum384(arg)
		sum = s[:]
	case 512:
		s := sha512.Sum512(arg)
		sum = s[:]
	}
	fmt.Printf("%x\n", sum)
}
