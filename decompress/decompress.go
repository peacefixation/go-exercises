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

	decompressed := Decompress([]rune("1[" + os.Args[1] + "]"))
	fmt.Println(string(decompressed))
}

// Decompress decompress a compressed string
func Decompress(text []rune) []rune {
	decompressed, _ := decompress(text, 0, 1)
	return decompressed
}

func decompress(text []rune, start, times int) ([]rune, int) {
	if times == 0 {
		return []rune{}, skip(text, start)
	}

	var decompressed []rune
	var end int
	for n := 0; n < times; n++ {
		i := start
		for ; i < len(text) && text[i] != ']'; i++ {
			if unicode.IsLetter(text[i]) {
				decompressed = append(decompressed, text[i])
			} else if unicode.IsDigit(text[i]) {
				subTimes := 0
				for unicode.IsDigit(text[i]) {
					subTimes = subTimes*10 + toInt(text[i])
					i++
				}
				i++ // iterate past the '['

				var chunk []rune
				chunk, i = decompress(text, i, subTimes)
				decompressed = append(decompressed, chunk...)
			}
		}
		end = i
	}
	return decompressed, end
}

// skip a chunk (the multiplier was zero)
func skip(text []rune, start int) int {
	i := start
	for ; i < len(text) && text[i] != ']'; i++ {
		if unicode.IsDigit(text[i]) {
			for ; unicode.IsDigit(text[i]); i++ {
				i++
			}
			i++ // iterate past the '['

			// skip the nested chunk
			i = skip(text, i)
		}
	}
	return i
}

// convert a rune to an integer, die if it fails
func toInt(r rune) int {
	n, err := strconv.Atoi(string(r))
	if err != nil {
		log.Fatalf("Conversion error: %v\n", err)
	}
	return n
}
