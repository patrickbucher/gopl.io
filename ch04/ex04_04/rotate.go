package ex04_04

// Rotate rotates slice s r times to the left.
func Rotate(s []int, r int) []int {
	n := len(s)
	rotated := make([]int, n)
	for i := range s {
		rotated[i] = s[(i+r)%n]
	}
	return rotated
}
