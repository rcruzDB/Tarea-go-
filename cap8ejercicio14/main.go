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

type client struct {
	name string
	ch   chan<- string 
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
				cli.ch <- msg
			}

		case cli := <-entering:
			clients[cli] = true

			
			var onlines []string
			for c := range clients {
				onlines = append(onlines, c.name)
			}
			cli.ch <- fmt.Sprintf("%d clients: %s", len(clients), strings.Join(onlines, ", "))

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	
	talk := make(chan struct{})

	ch := make(chan string) 
	go clientWriter(conn, ch)
	input := bufio.NewScanner(conn)

	
	var who string
	go func() {
		ch <- "Input your name:"
		if input.Scan() {
			who = input.Text()
			talk <- struct{}{}
		} else {
			leaving <- client{who, ch}
			messages <- who + " has left"
			conn.Close()
			return
		}
	}()

	
loop:
	for {
		select {
		case _, ok := <-talk:
			if ok {
				break loop
			} else {
				conn.Close()
				return
			}
		case <-time.After(timeout):
			conn.Close()
			return
		}
	}

	messages <- who + " has arrived"
	entering <- client{who, ch}

	go func() {
		for {
			if input.Scan() {
				messages <- who + ": " + input.Text()
				talk <- struct{}{}
			} else {
				leaving <- client{who, ch}
				messages <- who + " has left"
				conn.Close()
				return
			}
		}
	}()
	
	for {
		select {
		case _, ok := <-talk:
			if !ok {
				conn.Close()
				return
			}
		case <-time.After(timeout):
			conn.Close()
			return
		}
	}
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