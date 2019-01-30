package countingwriter

import (
	"bytes"
	"testing"
)

var tests = []struct {
	addition      string
	bufferedBytes int64
	bufferedValue string
}{
	{"abc", 3, "abc"},
	{"def", 6, "abcdef"},
	{"ghi", 9, "abcdefghi"},
	{"", 9, "abcdefghi"},
}

func TestCountingWriter(t *testing.T) {
	buf := bytes.NewBufferString("")
	writer, n := CountingWriter(buf)
	for i := range tests {
		add := tests[i].addition
		writer.Write([]byte(add))
		expectedBytes := tests[i].bufferedBytes
		if *n != expectedBytes {
			t.Errorf("added %q, expected length %d, was %d", add, expectedBytes, *n)
		}
		expectedString := tests[i].bufferedValue
		gotString := buf.String()
		if gotString != expectedString {
			t.Errorf("added %q, expected value %q, was %q", add, expectedString, gotString)
		}
	}
}
