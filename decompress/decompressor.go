package main

import (
	"log"
	"strconv"
	"unicode"
)

// Decompressor decompress compressed strings
// 3[abc]4[ab]c -> abcabcabcababababc
type Decompressor struct{}

// Decompress decompress a compressed string
func (d Decompressor) Decompress(text []rune) []rune {
	decompressed := []rune{}
	for r := range decompress(text, 0, 1) {
		decompressed = append(decompressed, r)
	}
	return decompressed
}

// decompress a string and return a channel that the characters can be read from
func decompress(text []rune, start, times int) <-chan rune {
	chnl := make(chan rune)
	go func() {
		defer close(chnl)
		decompressChunk(text, start, times, chnl)
	}()
	return chnl
}

// decompress a chunk of a string from a start index some number of times
// send the characters to the provided channel and return the index of the end of the 'chunk'
func decompressChunk(text []rune, start, times int, chnl chan rune) int {
	if times == 0 {
		return skip(text, start)
	}

	var end int
	for n := 0; n < times; n++ {
		i := start
		for ; i < len(text) && text[i] != ']'; i++ {
			if unicode.IsLetter(text[i]) {
				chnl <- text[i]
			} else if unicode.IsDigit(text[i]) {
				subTimes := 0
				for ; unicode.IsDigit(text[i]); i++ {
					subTimes = subTimes*10 + toInt(text[i])
				}
				i++ // iterate past the '['

				i = decompressChunk(text, i, subTimes, chnl)
			}
		}
		end = i
	}
	return end
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
