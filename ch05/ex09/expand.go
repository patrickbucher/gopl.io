package ex09

import "strings"

func Expand(s string, f func(string) string) string {
	return expand(s, f)
}

func expand(s string, f func(string) string) string {
	words := strings.Fields(s)
	var vars []string
	for _, w := range words {
		if w[0] == '$' {
			if len(w[1:]) > 0 {
				vars = append(vars, w)
			}
		}
	}
	for _, v := range vars {
		s = strings.Replace(s, v, f(v[1:]), -1)
	}
	return s
}
