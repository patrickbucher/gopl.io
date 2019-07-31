package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopl.io/ch07/eval"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("expression (such as 2*x + 3*y):\t")
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
		fmt.Println("variables (such as x=3, y=2):\t")
		input, err = stdin.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading string from stdin: %v\n", err)
			continue
		}
		definitions := strings.Split(input, ",")
		environment := make(map[eval.Var]float64)
		for _, def := range definitions {
			parts := strings.Split(def, "=")
			if len(parts) != 2 {
				fmt.Fprintf(os.Stderr, "parsing definition '%s' failed\n", parts)
				continue
			}
			// TODO: trim all parts
			key := eval.Var(parts[0])
			value, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "parse '%s' as float: %v\n", parts[1], err)
				continue
			}
			environment[key] = value
		}
		result := expr.Eval(eval.Env(environment))
		fmt.Println(result)
	}
}
