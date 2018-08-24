package ex05

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var pc int
	for ; x > 0; pc++ {
		x = x & (x - 1)
	}
	return pc
}
