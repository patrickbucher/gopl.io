package main

import "fmt"

const (
	KB float64 = 1000
	MB float64 = KB * 1000
	GB float64 = MB * 1000
	TB float64 = GB * 1000
	PB float64 = TB * 1000
	EB float64 = PB * 1000
	ZB float64 = EB * 1000
	YB float64 = ZB * 1000
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(ZB)
	fmt.Println(YB)
}
