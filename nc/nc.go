package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func usage() {
	fmt.Println("Usage: nc [options]")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	host := flag.String("h", "localhost", "the address to connect to")
	listen := flag.Bool("l", false, "the port to listen on")
	port := flag.Int("p", 1234, "the port to connect to")
	flag.Parse()

	protocol := "tcp"

	if *listen {
		var listener Listener
		err := listener.Listen(protocol, *host, *port)
		if err != nil {
			log.Fatalf("Listener error: %v\n", err)
		}
	} else {
		var client Client
		message := "some message"
		err := client.Send(protocol, *host, *port, message)
		if err != nil {
			log.Fatalf("Client error: %v\n", err)
		}
	}
}
