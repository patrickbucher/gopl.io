package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopl.io/ch04/ex11"
)

func main() {
	file, err := ioutil.TempFile("", "ex4.11")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = ex11.Edit(file.Name()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}
