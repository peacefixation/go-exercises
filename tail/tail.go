package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func usage() {
	fmt.Println("Usage: tail [-n x] file")
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
	} else {
		usage()
	}
}

// read (up to) numLines from the end of a file
func readLines(file *os.File, numLines int) {
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("File error: %v\n", err)
	}

	fileSize := fileInfo.Size()

	if fileSize == 0 {
		return
	}

	var offset int64 = -1 // start at (negative) offset of 1 byte
	var startPos int64
	var foundLines int

	for foundLines < numLines && abs(offset) <= fileSize {
		// seek a position offset from the end of the file
		startPos, err = file.Seek(offset, 2)
		if err != nil {
			log.Fatalf("Seek error: %v\n", err)
		}

		// read the byte at the offset
		b := make([]byte, 1)
		_, err = file.ReadAt(b, startPos)
		if err != nil {
			log.Fatalf("Read error: %v\n", err)
		}

		// if it's a newline, increment the counter
		if string(b) == "\n" {
			foundLines++
		}

		// increase the offset
		offset--
	}

	var bufSize int64 = 512

	// read the file in chunks of bufSize from the startPos to the end
	for {
		buf := make([]byte, bufSize)
		_, err := file.ReadAt(buf, startPos)
		if err == io.EOF {
			fmt.Print(string(buf))
			break
		} else if err != nil {
			log.Fatalf("Read error: %v\n", err)
		}

		fmt.Print(string(buf))
		startPos += bufSize
	}
}

// return the absolute value of an integer
func abs(n int64) int64 {
	if n < 0 {
		n = -n
	}
	return n
}
