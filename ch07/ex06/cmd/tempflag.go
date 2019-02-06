package main

import (
	"flag"
	"fmt"

	"gopl.io/ch07/ex06"
)

var temp = ex06.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
