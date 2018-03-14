package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func usage() {
	fmt.Println("Usage: wget [-o] <url>")
	os.Exit(1)
}

// https://golangcode.com/download-a-file-from-a-url/

func main() {
	outputFilenameParam := flag.String("o", "", "write the response body to a file")
	flag.Parse()

	if flag.NArg() != 1 {
		usage()
	}

	urlParam := flag.Arg(0)

	outputFilename := *outputFilenameParam
	if outputFilename == "" {
		outputFilename = determineOutputFilename(urlParam)
	}

	out, err := os.Create(outputFilename)
	if err != nil {
		fmt.Printf("Output file error: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()

	resp, err := http.Get(urlParam)
	if err != nil {
		fmt.Printf("HTTP GET error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP Status: %s\n", resp.Status)
		os.Exit(1)
	}

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("Copy error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Copied %d bytes to %s\n", n, outputFilename)
}

func determineOutputFilename(urlParam string) string {
	outputFilename := "index.html"

	url, err := url.Parse(urlParam)
	if err != nil {
		fmt.Printf("URL parse error: %v\n", err)
		os.Exit(1)
	}

	path := url.Path
	slashIndex := strings.LastIndex(path, "/")
	if slashIndex != -1 && slashIndex < len(path) {
		pathFilename := path[slashIndex+1:]
		if pathFilename != "" {
			outputFilename = pathFilename
		}
	}

	return outputFilename
}
