package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"bytes"
	"flag"
)

func usage() {
	fmt.Println("Usage: tr string1 string2")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	findChars := os.Args[1]
	replChars := os.Args[2]

	// read lines from stdin
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}

		output := substitute(input, findChars, replChars)
		fmt.Print(output)
	}
}

func substitute(input, findChars, replChars string) string {
	var buffer bytes.Buffer

	for _, char := range input {
		substituted := false
		for pos, c := range findChars {
			if c == char {
				buffer.WriteRune([]rune(replChars)[pos]);
				substituted = true
			}
		}

		if !substituted {
			buffer.WriteRune(char)
		}
	}

	return buffer.String()
}