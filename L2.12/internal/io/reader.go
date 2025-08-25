package io

import (
	"bufio"
	"os"
)

// ReadLines читает все строки из файла (или STDIN) и возвращает как []string.
func ReadLines(path string) ([]string, error) {
	var file *os.File
	var err error

	if path == "" {
		file = os.Stdin
	} else {
		file, err = os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
