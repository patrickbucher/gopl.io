package limitreader

import "io"

type limitReader struct {
	reader io.Reader
	limit  int64
	read   int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{reader: r, limit: n}
}

func (l *limitReader) Read(p []byte) (n int, err error) {
	left := l.limit - l.read
	if left == 0 {
		return 0, io.EOF
	}
	if int64(len(p)) < left {
		n, err = l.reader.Read(p)
	} else {
		n, err = l.reader.Read(p[:left])
		if err == nil {
			err = io.EOF
		}
	}
	l.read += int64(n)
	return n, err
}
