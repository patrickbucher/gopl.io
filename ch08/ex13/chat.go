package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	c    chan<- string // an outgoing message channel
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.c <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			for client := range clients {
				if client != cli {
					cli.c <- fmt.Sprintf("%s is also on the chat", client.name)
				}
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.c)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client{ch, who}

	reset := make(chan struct{})
	go func() {
		input := bufio.NewScanner(conn)
		for input.Scan() {
			messages <- who + ": " + input.Text()
			reset <- struct{}{}
		}
	}()

	timeout := 15 * time.Second
	disconnect := time.NewTimer(timeout)
loop:
	for {
		select {
		case <-disconnect.C:
			log.Println("close connection to", who, "after", timeout, "of inactivity")
			conn.Close()
			break loop
		case <-reset:
			log.Println("reset timeout of", who, "to", timeout, "after activity")
			disconnect.Reset(timeout)
		}
	}

	// NOTE: ingoring potential errors from input.Err()
	leaving <- client{ch, who}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
