package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"
)

func usage() {
	fmt.Println("Usage: wc [-l] [-w] [-c] file")
	os.Exit(1)
}

func main() {

	showLines := flag.Bool("l", false, "show the number of lines in the input file")
	showWords := flag.Bool("w", false, "show the number of words in the input file")
	showChars := flag.Bool("c", false, "show the number of characters in the input file")
	flag.Parse()

	if *showLines == false && *showWords == false && *showChars == false {
		// if no flags are supplied, default to show all
		*showLines = true
		*showWords = true
		*showChars = true
	}

	numLines := 0
	numWords := 0
	numChars := 0

	info, _ := os.Stdin.Stat()

	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		// filename as argument
		if flag.NArg() != 1 {
			usage()
		}
	
		filename := flag.Arg(0);

		file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
		if err != nil {
			exit("File error", err)
		}
		defer file.Close()

		numLines, numWords, numChars = readFile(file, *showWords, *showChars)
	} else if info.Size() > 0 {
		// piped input, read lines from stdin
		numLines, numWords, numChars = readFile(os.Stdin, *showWords, *showChars)
	}

	totals := formatTotal(numLines, *showLines) + formatTotal(numWords, *showWords) + formatTotal(numChars, *showChars)
	fmt.Println(totals)
}

// read a file and count the lines (and words, and chars if specified)
func readFile(file *os.File, showWords, showChars bool) (int, int, int) {
	numLines := 0
	numWords := 0
	numChars := 0

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			exit("Read error", err)
		}

		numLines++
		
		if showWords {
			numWordsInLine, numCharsInLine := countWords(line)
			numWords += numWordsInLine
			numChars += numCharsInLine
		} else if showChars {
			numChars += countChars(line)
		}
	}

	return numLines, numWords, numChars
}

// count the number of words (and characters) in a line
func countWords(line string) (int, int) {
	numWords := 0
	numChars := 0
	isWord := false

	for _, r := range line {
		if unicode.IsSpace(r) && isWord {
			numWords++
			isWord = false
		} else if !unicode.IsSpace(r) && !isWord {
			isWord = true
		}

		numChars++
	}

	return numWords, numChars
}

// count the number of characters in a line
func countChars(line string) int {
	return len([]rune(line))
}

// print the total if the value of show is true
func formatTotal(num int, show bool) string {
	if show {
		return fmt.Sprintf("%8d", num)
	} else {
		return ""
	}
}

// print an error message and exit with status 1
func exit(message string, err error) {
	fmt.Printf("%s: %v\n", message, err)
	os.Exit(1)
}
