package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var usage = func() {
	fmt.Printf("Usage: %s [options] file1 file2 ...\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	numberLines := flag.Bool("n", false, "number output lines")
	flag.Parse()

	info, _ := os.Stdin.Stat()

	startLineNumber := -1
	if *numberLines {
		startLineNumber = 1
	}

	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		// filename as argument
		if flag.NArg() < 1 {
			usage()
		}

		// print each file to stdout
		for _, filename := range flag.Args() {
			file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
			if err != nil {
				log.Fatalf("File error: %v\n", err)
			}
			defer file.Close()
			startLineNumber = printFile(file, startLineNumber)
		}
	} else if info.Size() > 0 {
		// read lines from stdin
		printFile(os.Stdin, startLineNumber)
	}
}

// print a file to stdout
func printFile(file *os.File, startLineNumber int) int {
	reader := bufio.NewReader(file)
	for {
		numberStr := ""
		if startLineNumber != -1 {
			numberStr = strconv.Itoa(startLineNumber) + "  "
		}

		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("%8s%s\n", numberStr, line)
			break
		}

		if err != nil {
			log.Fatalf("Read error: %v\n", err)
		}

		fmt.Printf("%8s%s", numberStr, line)
		startLineNumber++
	}

	startLineNumber++ // increment for next file
	return startLineNumber
}
