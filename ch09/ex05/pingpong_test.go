package pingpong

import (
	"testing"
	"time"
)

func TestPingpongMutex(t *testing.T) {
	n := PingpongMutex("hello", time.Second)
	t.Log(n)
}
