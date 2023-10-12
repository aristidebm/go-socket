package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	// "time"
)

var wg = sync.WaitGroup{}

func main() {
	Client()
}

func Client() {
	netFamily := "tcp"
	address := "127.0.0.1:8000"
	conn, err := net.Dial(netFamily, address)

	if err != nil {
		log.Fatalf("Cannot connect to %s", address)
	}

	fmt.Printf("(%s)> ", conn.LocalAddr())
	msg, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		log.Fatal("Cannot read your message")
	}

	defer conn.Close()

	msg = fmt.Sprintf("(%s)> %s", conn.LocalAddr(), msg)

	fmt.Fprint(conn, msg)
	// _, err = conn.Write([]byte(msg))

	// time.Sleep(time.Second * 4)

	// for {
	// fmt.Fprint(conn, msg)
	wg.Add(1)
	go HandleResponse(conn, &wg)
	wg.Wait()
	// }

}

func HandleResponse(conn net.Conn, wg *sync.WaitGroup) {
	r := bufio.NewReader(conn)
	body, err := r.ReadString('\n')

	if err != nil && err != io.EOF {
		log.Printf("Sorry I can't here you %s", conn.RemoteAddr())
		wg.Done()
		return
	}

	fmt.Println(body)
	wg.Done()
}
