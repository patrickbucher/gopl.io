package main

import "fmt"

func main() {
	fmt.Println(sum())           // 0
	fmt.Println(sum(1))          // 1
	fmt.Println(sum(1, 2))       // 3
	fmt.Println(sum(1, 2, 3))    // 6
	fmt.Println(sum(1, 2, 3, 4)) // 10

	numbers := []int{1, 2, 3, 4}
	fmt.Println(sum(numbers...)) // 10
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
