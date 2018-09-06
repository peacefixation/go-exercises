package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Client send a network request
type Client struct {
}

// Send send a network request
func (c Client) Send(protocol, ip string, port int, message string) error {
	connection, err := net.Dial(protocol, fmt.Sprintf("%s:%d", ip, port))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Connection error: %v\n", err)
		return err
	}

	defer connection.Close()

	connection.Write([]byte(message))

	handleResponse(connection)

	return nil
}

func handleResponse(connection net.Conn) {
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
}
