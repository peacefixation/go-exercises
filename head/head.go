package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func usage() {
	fmt.Println("Usage: head [-n x] file")
	os.Exit(1)
}

func main() {
	numLines := flag.Int("n", 10, "the number of lines")
	flag.Parse()

	info, _ := os.Stdin.Stat()

	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		// filename as argument
		if flag.NArg() != 1 {
			usage()
		}

		filename := flag.Arg(0)

		file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Fatalf("File error: %v\n", err)
		}
		defer file.Close()

		readLines(file, *numLines)
	} else if info.Size() > 0 {
		// read lines from stdin
		readLines(os.Stdin, *numLines)
	}
}

// read (up to) numLines from the start of a file
func readLines(file *os.File, numLines int) {
	var foundLines int

	reader := bufio.NewReader(file)
	for foundLines < numLines {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print(line)
			break
		}

		if err != nil {
			log.Fatalf("Read error: %v\n", err)
		}

		fmt.Print(line)
		foundLines++
	}
}
