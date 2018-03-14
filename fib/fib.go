package main

import "fmt"
import "os"
import "strconv"

func usage() {
	fmt.Println("usage: fib n (1-92)")
	os.Exit(1)
}

var memo = map[uint64]uint64{}

func fib(n uint64) uint64 {
	if n == 1 || n == 2 {
		return 1
	}

	fibN, exists := memo[n]
	if exists {
		return fibN
	}

	memo[n] = fib(n - 1) + fib(n - 2)
	return memo[n]
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	number, err := strconv.ParseUint(os.Args[1], 10, 64)
	if(err != nil) {
		usage()
	}

	if(number < 1) {
		usage()
	}

	if(number > 92) {
		usage()
	}

	fibN := fib(number)
	
	fmt.Printf("%d\n", fibN)
}
