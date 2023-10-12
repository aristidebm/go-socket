package main

import (
	// "bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	Server()
}

func Server() {
	netFamily := "tcp"
	address := "127.0.0.1:8000"
	listener, err := net.Listen(netFamily, address)
	defer listener.Close()

	if err != nil {
		log.Fatalf("Cannot open a connection to %s\n", address)
	}

	log.Printf("The server is started and is listening at %s\n", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Cannot process a request coming from %s\n", conn.RemoteAddr())
			continue
		}
		log.Printf("Mr %s is entering the chatroom\n", conn.RemoteAddr())
		go HandleRequest(conn)
	}
}

func HandleRequest(conn net.Conn) {
	defer conn.Close()
	log.Printf("Processing the message of Mr %s", conn.RemoteAddr())
	resp := fmt.Sprintf("(%s)> Hello Mr %s, you are welcome in the chat room", conn.LocalAddr(), conn.RemoteAddr())
	// Note: For a reason that I don't understand yet
	// using bufio Readers here, hangs the request body
	// till the client cloe de connexion
	// r := bufio.NewReader(conn)
	// body, err := r.ReadString('\n')
	body := make([]byte, 1024)
	_, err := conn.Read(body)

	if err != nil && err != io.EOF {
		log.Printf("Sorry, we cannot understand Mr %s\n", conn.RemoteAddr())
		return
	}

	fmt.Println(string(body))

	_, err = conn.Write([]byte(resp))
	if err != nil {
		fmt.Printf("Cannot respond to %s\n", conn.RemoteAddr())
	}
}
