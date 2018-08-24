package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch02/lenconv"
	"gopl.io/ch02/tempconv"
	"gopl.io/ch02/weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		f, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse %v to float: %v\n", arg, err)
		}
		fmt.Printf("%s = %s, %s = %s\n",
			tempconv.Fahrenheit(f), tempconv.FToC(tempconv.Fahrenheit(f)),
			tempconv.Celsius(f), tempconv.CToF(tempconv.Celsius(f)))
		fmt.Printf("%s = %s, %s = %s\n",
			weightconv.Pound(f), weightconv.PToK(weightconv.Pound(f)),
			weightconv.Kilogram(f), weightconv.KToP(weightconv.Kilogram(f)))
		fmt.Printf("%s = %s, %s = %s\n",
			lenconv.Foot(f), lenconv.FToM(lenconv.Foot(f)),
			lenconv.Meter(f), lenconv.MToF(lenconv.Meter(f)))
	}
}
