package main

import "fmt"

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println(p)
		}
	}()
	sumByPanicking(19, 23)
}

func sumByPanicking(a, b int) {
	defer func() {
		c := a + b
		panic(c)
	}()
}
