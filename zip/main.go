package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// TODO register compressor
// https://golang.org/src/archive/zip/example_test.go

// TODO get the compressed size of each file
// https://groups.google.com/forum/#!topic/Golang-Nuts/8AZ3wfseoJE

func main() {
	err := zipr("/Users/matthew/go/src/exercises/zip", "/Users/matthew/Desktop/files.zip")
	if err != nil {
		log.Fatal(err)
	}
}

// zip files at the path recursively
func zipr(pathToZip, destinationPath string) error {
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}

	zipWriter := zip.NewWriter(destinationFile)

	err = filepath.Walk(pathToZip, func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if err != nil {
			return err
		}

		relPath := strings.TrimPrefix(filePath, pathToZip)
		relPath = strings.TrimPrefix(relPath, "/") // remove the slash if it exists

		fmt.Print("  adding " + relPath)

		fileToZip, err := os.Open(filePath)
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = relPath       // preserve the relative path
		header.Method = zip.Deflate // compress the file

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		written, err := io.Copy(writer, fileToZip)
		if err != nil {
			return err
		}

		// TODO print the compression ratio (original size / compressed size)
		compressedSize := 1.0                                 // TODO calculate
		compressionRatio := compressedSize / float64(written) // TODO handle file size 0
		fmt.Printf(" (compression %0.f%%)\n", compressionRatio)

		return nil
	})

	if err != nil {
		return err
	}

	err = zipWriter.Close()
	if err != nil {
		return err
	}

	return nil
}
