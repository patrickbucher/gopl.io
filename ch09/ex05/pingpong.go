package pingpong

import (
	"sync"
	"time"
)

func PingpongMutex(message string, duration time.Duration) uint {
	var (
		mutex   sync.Mutex
		counter uint
	)
	pingpong := func(source <-chan string, target chan<- string, done <-chan struct{}) {
		for {
			select {
			case msg := <-source:
				target <- msg
				mutex.Lock()
				counter++
				mutex.Unlock()
			case <-done:
				close(target)
				return
			}
		}
	}

	a, b := make(chan string), make(chan string)
	done := make(chan struct{}, 2)

	go pingpong(a, b, done)
	go pingpong(b, a, done)
	go func() { a <- message }()
	finished := time.After(duration)

	<-finished
	done <- struct{}{} // one for ping...
	done <- struct{}{} // ... and one for pong

	defer mutex.Unlock()
	mutex.Lock()
	return counter
}
