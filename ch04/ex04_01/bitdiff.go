package ex04_01

func BitDiff(a, b []byte) int {
	if len(a) != len(b) {
		return -1
	}
	var diff int
	for i := 0; i < len(a); i++ {
		c := a[i] ^ b[i]
		diff += int(PopCount(c))
	}
	return diff
}

func PopCount(b byte) uint8 {
	var pop uint8
	for i := uint8(0); i < uint8(8); i++ {
		if b&1 == 1 {
			pop++
		}
		b = b >> 1
	}
	return pop
}
