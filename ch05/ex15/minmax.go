package ex15

import "math"

func Min(numbers ...int64) (min int64, found bool) {
	if len(numbers) < 1 {
		return 0, false
	}
	min = math.MaxInt64
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}
	return min, true
}

func Max(numbers ...int64) (max int64, found bool) {
	if len(numbers) < 1 {
		return 0, false
	}
	max = math.MinInt64
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max, true
}
