package main

import "fmt"
import "os"
import "strconv"

func usage() {
	fmt.Println("usage: avg n1 n2 n3 ...")
	os.Exit(1)
}

func parseArgs(args []string) []float64 {
	numbers := make([]float64, len(args) - 1)
	for i, arg := range args[1:] {
		number, err := strconv.ParseFloat(arg, 64)
		if(err != nil) {
			usage()
		}
		numbers[i] = number
	}
	return numbers
}

func avg(xs []float64) float64 {
	sum := 0.0
	for _, x := range xs {
		sum += x
	}
	return sum / float64(len(xs))
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	numbers := parseArgs(os.Args)

	avg := avg(numbers)
	
	fmt.Printf("%0.1f\n", avg)
}
