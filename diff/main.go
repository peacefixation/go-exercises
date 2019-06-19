package main

import (
	"exercises/diff/file"
	"exercises/diff/lcs"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	if len(flag.Args()) != 2 {
		fmt.Println("Must provide 2 filenames as arguments")
		os.Exit(1)
	}

	files := make([][]string, 0)

	for _, filename := range flag.Args() {
		file, err := file.Read(filename)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		files = append(files, file)
	}

	xs := files[0]
	ys := files[1]

	c := lcs.ComputeLCS(xs, ys)
	lcs.PrintDiffRecursive(c, xs, ys, len(xs), len(ys))
}
