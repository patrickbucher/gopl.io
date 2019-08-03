package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type timeServer struct {
	location string
	host     string
	port     uint16
}

func (ts *timeServer) String() string {
	return fmt.Sprintf("%s=%s:%d", ts.location, ts.host, ts.port)
}

func (ts *timeServer) Socket() string {
	return fmt.Sprintf("%s:%d", ts.host, ts.port)
}

func parseTimeServer(spec string) (*timeServer, error) {
	var ts timeServer
	parts := strings.Split(spec, "=")
	if len(parts) != 2 {
		err := fmt.Errorf("spec '%s' malformed: [location]=[host]:[port]", spec)
		return nil, err
	}
	ts.location = parts[0]
	serverParts := strings.Split(parts[1], ":")
	if len(serverParts) != 2 {
		err := fmt.Errorf("server '%s' malformed: [host]:[port]", serverParts)
		return nil, err
	}
	ts.host = serverParts[0]
	port, err := strconv.Atoi(serverParts[1])
	if err != nil {
		return nil, fmt.Errorf("parse '%s' as port number: %v", serverParts[1], err)
	}
	if port < 0 || port > 65535 {
		return nil, fmt.Errorf("%d is not a port number (0..65536)", port)
	}
	ts.port = uint16(port)
	return &ts, nil
}

type timeInfo struct {
	location string
	time     string
}

func main() {
	var servers []*timeServer
	for _, arg := range os.Args[1:] {
		server, err := parseTimeServer(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse %s as time server spec: %v\n", arg, err)
			continue
		}
		servers = append(servers, server)
	}
	if len(servers) == 0 {
		fmt.Fprintln(os.Stderr, "no suitable time servers defined, exiting")
		os.Exit(1)
	}
	timeChan := make(chan timeInfo)
	times := make(map[string]string)
	go func() {
		for {
			select {
			case msg := <-timeChan:
				times[msg.location] = msg.time
			}
		}
	}()
	for _, server := range servers {
		conn, err := net.Dial("tcp", server.Socket())
		if err != nil {
			fmt.Fprintf(os.Stderr, "connecting to %s failed: %v", server.Socket(), err)
			continue
		}
		defer conn.Close()
		go func(c chan<- timeInfo, s *timeServer, r io.Reader) {
			buf := make([]byte, 100)
			for {
				n, err := r.Read(buf)
				if err != nil {
					if err == io.EOF {
						return
					}
					fmt.Fprintf(os.Stderr, "reading from %s: %v", s, err)
				}
				if n > 0 {
					time := strings.TrimSpace(string(buf[:n]))
					c <- timeInfo{s.location, time}
				} else {
					fmt.Fprintf(os.Stderr, "nothing read from %s", s)
				}
			}
		}(timeChan, server, conn)
	}
	output(times)
}

func output(times map[string]string) {
	for {
		var keys []string
		for k := range times {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Printf("%s: %s\t", k, times[k])
		}
		fmt.Println()
		time.Sleep(time.Second)
	}
}
