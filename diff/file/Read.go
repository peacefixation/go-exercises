package file

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Read read the file with the given filename
func Read(filename string) ([]string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([]string, 0)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		lines = append(lines, strings.TrimSuffix(line, "\n"))
	}

	return lines, nil
}
