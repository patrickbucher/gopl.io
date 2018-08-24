package echo

import (
	"fmt"
	"strings"
)

func Echo1(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func Echo2(args []string) {
	var s, sep string
	for _, arg := range args {
		s += sep + arg
		s = " "
	}
	fmt.Println(s)
}

func Echo3(args []string) {
	fmt.Println(strings.Join(args, " "))
}
