package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopl.io/ch04/ex11"
)

func main() {
	var file *os.File
	var data []byte
	if f, err := ioutil.TempFile("", "ex4.11"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		file = f
	}
	if err := ex11.Edit(file.Name()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if d, err := ioutil.ReadFile(file.Name()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		data = d
	}
	fmt.Println(string(data))
}
