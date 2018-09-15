package ex16

import "bytes"

func Join(sep string, parts ...string) string {
	if len(parts) == 0 {
		return ""
	}
	buf := bytes.NewBufferString(parts[0])
	for _, p := range parts[1:] {
		buf.WriteString(sep)
		buf.WriteString(p)
	}
	return buf.String()
}
