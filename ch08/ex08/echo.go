package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	ticker := time.NewTicker(10 * time.Second)
	alive := make(chan struct{})
	go func(alive chan<- struct{}) {
		input := bufio.NewScanner(c)
		for input.Scan() {
			alive <- struct{}{}
			go echo(c, input.Text(), 1*time.Second)
		}
	}(alive)
	// NOTE: ignoring potential errors from input.Err()
	for {
		select {
		case <-ticker.C:
			log.Println("closing connection after 10 seconds of inactivity")
			c.Close()
			ticker.Stop()
		case <-alive:
			log.Println("keeping connection open for another 10 seconds after activity")
			ticker.Stop()
			ticker = time.NewTicker(10 * time.Second)
		}
	}
}
