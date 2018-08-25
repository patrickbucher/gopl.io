package ex05

func Uniq(s []string) []string {
	if len(s) < 2 {
		return s
	}
	var r, w int
	for r, w = 1, 0; r < len(s); r++ {
		if s[r] != s[w] {
			w++
			s[w] = s[r]
		}
	}
	return s[:w+1]
}
