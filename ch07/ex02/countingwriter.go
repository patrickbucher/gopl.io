package countingwriter

import "io"

// Wraps a counting Writer around the given Writer and returns the new wrapping
// Writer alongside a pointer to its byte counter.
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var counter Counter
	counter.writer = w
	return &counter, &counter.count
}

type Counter struct {
	count  int64
	writer io.Writer
}

func (c *Counter) Write(p []byte) (n int, err error) {
	n, err = c.writer.Write(p)
	c.count += int64(n)
	return n, err
}
