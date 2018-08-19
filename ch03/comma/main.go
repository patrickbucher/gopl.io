package main

import "fmt"

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	input := []int{
		233,
		3234,
		876324,
		959743928,
		37237489239,
	}
	for i := range input {
		s := fmt.Sprintf("%d", input[i])
		fmt.Printf("%12d => %14s\n", input[i], comma(s))
	}
}
