package ex12

import (
	"strings"
)

func IsAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for _, r := range a {
		if i := strings.Index(b, string(r)); i == -1 {
			return false
		} else {
			b = b[:i] + b[i+1:]
		}
	}
	return true
}
