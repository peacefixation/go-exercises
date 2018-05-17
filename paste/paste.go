package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var usage = func() {
	fmt.Printf("Usage: %s [options] file1 file2 ...\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	delimiter := flag.String("d", "\t", "delimiter")
	flag.Parse()

	info, _ := os.Stdin.Stat()

	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		// filenames as arguments
		if flag.NArg() < 1 {
			usage()
		}

		files := make([]*os.File, flag.NArg())

		for i, filename := range flag.Args() {
			file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
			if err != nil {
				log.Fatalf("File error: %v\n", err)
			}
			defer file.Close()

			files[i] = file
		}

		concatLines(files, *delimiter)
	} else {
		usage()
	}
}

//
func concatLines(files []*os.File, delimiter string) {
	readers := make([]*bufio.Reader, len(files))
	for i, file := range files {
		readers[i] = bufio.NewReader(file)
	}

	for {
		lineElements := make([]string, len(readers))
		readErrors := make([]error, len(readers)) // keep track of reader errors (io.EOF)

		for i, reader := range readers {
			if readErrors[i] == nil {
				lineElements[i], readErrors[i] = readLine(reader)
			}
		}

		// if any files are still open, or we have a non-blank line
		if anyNil(readErrors) || !allBlank(lineElements) {
			fmt.Println(strings.Join(lineElements, delimiter))
		} else {
			break
		}
	}
}

// read a line
func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	return line, err
}

// check if the array only contains blank strings
func allBlank(strs []string) bool {
	for _, str := range strs {
		if str != "" {
			return false
		}
	}
	return true
}

// check of any of the error array elements are nil
func anyNil(readErrors []error) bool {
	for _, err := range readErrors {
		if err == nil {
			return false
		}
	}
	return true
}
