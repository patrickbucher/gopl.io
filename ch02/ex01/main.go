package main

import (
	"fmt"

	"gopl.io/ch02/tempconv"
)

func main() {
	c := tempconv.Celsius(32.5)
	f := tempconv.CToF(c)
	k := tempconv.CToK(c)
	fmt.Printf("%v = %v = %v\n", c, f, k)

	f = 100
	c = tempconv.FToC(f)
	k = tempconv.FToK(f)
	fmt.Printf("%v = %v = %v\n", c, f, k)

	k = 300
	f = tempconv.KToF(k)
	c = tempconv.KToC(k)
	fmt.Printf("%v = %v = %v\n", c, f, k)
}
