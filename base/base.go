package main

import (
	"fmt"
	"os"
	"strconv"
)

func usage() {
	fmt.Println("usage: base n b")
	os.Exit(1)
}

func parseInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Failed to convert '%s' to an integer\n", str)
		os.Exit(1)
	}

	return num
}

func convert(number, base int) int {
	result := ""
	for number > 0 {
		result = strconv.Itoa(number % base) + result;
		number /= base;
	}

	return parseInt(result)
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	number := parseInt(os.Args[1])
	base := parseInt(os.Args[2])

	converted := convert(number, base)
	
	fmt.Printf("%d\n", converted)
}
