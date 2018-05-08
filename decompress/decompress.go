package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func usage() {
	fmt.Println("usage: decompress string")
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	decompressed := Decompress([]rune(os.Args[1]))
	fmt.Println(string(decompressed))
}

// Decompress decompresses a compressed text
func Decompress(text []rune) []rune {
	length := len(text)
	var decompressed []rune
	for i := 0; i < length; i++ {
		for ; i < length && unicode.IsLetter(text[i]); i++ {
			decompressed = append(decompressed, text[i])
		}

		if i < length && unicode.IsDigit(text[i]) {
			var chunk []rune
			chunk, i = decompressChunk(text, i) // increment to the end of the chunk
			decompressed = append(decompressed, chunk...)
		}
	}
	return decompressed
}

func decompressChunk(text []rune, start int) ([]rune, int) {
	multiplier, startChunk := multiplier(text, start) // determine how many times to expand the chunk, and the starting point

	// edge case, skip past a chunk if the multiplier is 0
	if multiplier == 0 {
		return []rune{}, endOfChunk(text, startChunk)
	}

	var decompressed []rune // the decompressed string
	var endChunk int        // the index at the end of the chunk
	for n := 0; n < multiplier; n++ {
		i := startChunk // start decompression at the beginning of the chunk
		for ; i < len(text) && text[i] != ']'; i++ {
			if unicode.IsLetter(text[i]) {
				decompressed = append(decompressed, text[i])
			} else if unicode.IsDigit(text[i]) {
				var chunk []rune
				chunk, i = decompressChunk(text, i) // increment to the end of the chunk
				decompressed = append(decompressed, chunk...)
			}
		}
		endChunk = i // we do this every loop, unfortunately
	}

	return decompressed, endChunk
}

// determine the chunk multiplier and the startOfChunk index (i.e. the multiplier for 31[aaa] is 31 and the index is 3)
func multiplier(text []rune, start int) (int, int) {
	multiplier := 0
	i := start
	for ; i < len(text) && unicode.IsDigit(text[i]); i++ {
		multiplier = multiplier*10 + toInt(text[i])
	}
	i++ // increment past the '['
	return multiplier, i
}

// skip to the end of a chunk (because the multiplier is 0)
func endOfChunk(text []rune, start int) int {
	i := start
	for ; i < len(text) && text[i] != ']'; i++ {
		// also skip any nested chunks
		if unicode.IsDigit(text[i]) {
			_, startChunk := multiplier(text, i)
			i = endOfChunk(text, startChunk)
		}
	}
	return i
}

// convert a rune to an integer, die if it fails, we should check before calling
func toInt(r rune) int {
	n, err := strconv.Atoi(string(r))
	if err != nil {
		log.Fatalf("Conversion error: %v\n", err)
	}
	return n
}
