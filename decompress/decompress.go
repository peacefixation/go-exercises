package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("usage: decompress string")
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	d := Decompressor{}
	decompressed := d.Decompress([]rune(os.Args[1]))
	fmt.Println(string(decompressed))
}
