package limitreader

import (
	"io"
	"strings"
	"testing"
)

const input = "123456789"

var tests = []struct {
	input        string
	limit        int64
	expectedRead int
	expectedErr  error
}{
	{"123456789", 5, 5, io.EOF},
	{"123456789", 9, 9, io.EOF},
	{"123456789", 20, 9, nil},
}

func TestLimitReader(t *testing.T) {
	for i := range tests {
		reader := strings.NewReader(tests[i].input)
		limit := tests[i].limit
		limitReader := LimitReader(reader, int64(limit))
		buffer := make([]byte, len(input))
		n, err := limitReader.Read(buffer)
		expectedRead := tests[i].expectedRead
		if n != expectedRead {
			t.Errorf("limit was %d but read %d\n", expectedRead, n)
		}
		expectedErr := tests[i].expectedErr
		if err != expectedErr {
			t.Errorf("%v was expected, but err was %v\n", expectedErr, err)
		}
	}
}
