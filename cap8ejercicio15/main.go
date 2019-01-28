package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const timeout = 5 * time.Minute
const outbuffer = 64 

type client struct {
	name  string
	inch  chan<- string 
	outch chan<- string 
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) 
)

func broadcaster() {
	clients := make(map[client]bool) 
	for {
		select {
		case msg := <-messages:
			
			
			for cli := range clients {
				cli.outch <- msg
			}

		case cli := <-entering:
			clients[cli] = true

			
			var onlines []string
			for c := range clients {
				onlines = append(onlines, c.name)
			}
			cli.outch <- fmt.Sprintf("%d clients: %s", len(clients), strings.Join(onlines, ", "))

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.outch)
		}
	}
}

func handleConn(conn net.Conn) {
	inch := make(chan string)
	outch := make(chan string, outbuffer)

	go clientReader(conn, inch)
	go clientWriter(conn, outch)

	
	var who string

	outch <- "Input your name:"

	
	select {
	case in, ok := <-inch:
		if !ok {
			conn.Close()
			return
		}
		who = in
	case <-time.After(timeout):
		conn.Close()
		return
	}

	messages <- who + " has arrived"
	entering <- client{who, inch, outch}

	for {
		select {
		case in, ok := <-inch:
			if ok {
				messages <- who + ": " + in
			} else {
				leaving <- client{who, inch, outch}
				messages <- who + " has left"
				conn.Close()
				return
			}
		case <-time.After(timeout):
			leaving <- client{who, inch, outch}
			messages <- who + " has left"
			conn.Close()
			return
		}
	}
}

func clientReader(conn net.Conn, ch chan<- string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ch <- input.Text()
	}
	close(ch)
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) 
	}
}

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