package main

import "fmt"
import "os"
import "strconv"

func usage() {
	fmt.Println("usage: ftoc temp")
}

func ftoc(f float64) float64 {
	return (f - 32.0) * (5.0 / 9.0)
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	fahrenheit, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		usage()
		os.Exit(1)
	}

	celcius := ftoc(fahrenheit)

	fmt.Printf("%0.2f\n", celcius)
}
