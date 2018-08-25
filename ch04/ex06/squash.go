package ex06

import "unicode"

func Squash(s string) string {
	var w, spaces int
	runes := []rune(s)
	for r := 0; r < len(runes); r++ {
		if unicode.IsSpace(runes[r]) {
			if spaces == 0 {
				runes[w] = ' '
				w++
			}
			spaces++
		} else {
			runes[w] = runes[r]
			w++
			spaces = 0
		}
	}
	return string(runes[:w])
}
