package ex07

import "bytes"

// Reverse reverses a byte array representing an unicode encoded string.
func Reverse(b []byte) []byte {
	// TODO: is this cheating? Am I supposed to work on a bit level?
	r := bytes.Runes(b)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return []byte(string(r))
}
