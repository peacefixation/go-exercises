package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Listener listen for network requests
type Listener struct {
}

// Listen listen for network requests
func (l Listener) Listen(protocol, ip string, port int) error {
	address := fmt.Sprintf("%s:%d", ip, port)
	listener, err := net.Listen(protocol, address)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Connection error: %v\n", err)
		return err
	}

	defer listener.Close()

	fmt.Fprintf(os.Stdout, "Listening on %s\n", address)

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Connection error: %v\n", err)
			return err
		}

		go handleRequest(connection)
	}
}

func handleRequest(connection net.Conn) {
	var buf bytes.Buffer
	tmp := make([]byte, 1024)

	for {
		n, err := connection.Read(tmp)
		if err != nil {
			if err != io.EOF {
				log.Fatalf("Read error: %v\n", err)
			}
			break
		}

		for i := 0; i < n; i++ {
			if tmp[i] == '\n' {
				fmt.Println(buf.String())
				buf.Reset()
			} else {
				buf.WriteByte(tmp[i])
			}
		}
	}

	connection.Write([]byte("Request received.\n"))
	connection.Close()
}
