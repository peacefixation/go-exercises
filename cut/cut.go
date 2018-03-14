package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func usage() {
	fmt.Println("Usage: cut [-delimiter d] [-fields n] text")
	os.Exit(1)
}

func main() {
	delimiter := flag.String("delimiter", ",", "the delimiter to split on")
	fields := flag.String("fields", "", "the fields to print")
	flag.Parse()

	fieldsArray := parseFields(*fields)

	info, _ := os.Stdin.Stat()

	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		// input argument, use first non-flag argument
		if flag.NArg() != 1 {
			usage()
		}
		printFields(flag.Arg(0), *delimiter, fieldsArray)
	} else if info.Size() > 0 {
		// piped input, read lines from stdin
		reader := bufio.NewReader(os.Stdin)
		for {
			input, err := reader.ReadString('\n')
			if err != nil && err == io.EOF {
				break
			}

			input = strings.TrimSuffix(input, "\n")

			printFields(input, *delimiter, fieldsArray)
		}
	}
}

/**
 * Parse the fields argument to an array of integers
 */
func parseFields(fields string) []int {
	if fields != "" {
		fieldStrs := strings.Split(fields, ",")
		fieldsArray := make([]int, len(fieldStrs))

		for i := 0; i < len(fieldStrs); i++ {
			// is the field numeric?
			f, err := strconv.Atoi(fieldStrs[i])
			if err != nil {
				usage()
			} else {
				fieldsArray[i] = f
			}
		}
		return fieldsArray
	} else {
		// no fields were specified
		return []int{}
	}
}

/**
 * Split the input string on the delimiter and print the specified fields.
 */
func printFields(input string, delimiter string, fields []int) {
	if len(fields) == 0 {
		// print all the fields, we don't need to split anything
		fmt.Println(input)
	} else {
		parts := strings.Split(input, delimiter)
		ordered := []string{}
		for i := 0; i < len(fields); i++ {
			field := fields[i] - 1
			if field < len(parts) {
				ordered = append(ordered, parts[field])
			}
		}
		fmt.Println(strings.Join(ordered, delimiter))
	}
}
