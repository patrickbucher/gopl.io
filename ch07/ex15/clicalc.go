package main

import (
	"bufio"
	"fmt"
	"os"

	"gopl.io/ch07/eval"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("expression (such as 2*x + 3*y):\t")
		input, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading string from stdin: %v\n", err)
			continue
		}
		expr, err := eval.Parse(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parsing input '%s': %v\n", input, err)
			continue
		}
		fmt.Printf("variables (such as x=3, y=2):\t")
		input, err = stdin.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading string from stdin: %v\n", err)
			continue
		}
		environment, err := eval.ParseVarDefs(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse '%s': %v", input, err)
			continue
		}
		result := expr.Eval(environment)
		fmt.Println(result)
	}
}
