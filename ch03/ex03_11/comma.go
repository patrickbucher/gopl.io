package ex03_10

import (
	"bytes"
	"strings"
)

func Comma(s string) string {
	var buf bytes.Buffer
	if s[0] == '+' || s[0] == '-' {
		buf.WriteRune(rune(s[0]))
		s = s[1:]
	}
	var pointAt = strings.Index(s, ".")
	if pointAt == -1 {
		buf.WriteString(integerPart(s))
		return buf.String()
	}
	buf.WriteString(integerPart(s[:pointAt]))
	buf.WriteRune(rune('.'))
	buf.WriteString(decimalPart(s[pointAt+1:]))
	return buf.String()
}

func integerPart(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if len(s[i:len(s)])%3 == 0 && buf.Len() > 0 {
			buf.WriteRune(rune(','))
		}
		buf.WriteRune(rune(s[i]))
	}
	return buf.String()
}

func decimalPart(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if i > 0 && i%3 == 0 {
			buf.WriteRune(rune(','))
		}
		buf.WriteRune(rune(s[i]))
	}
	return buf.String()
}
