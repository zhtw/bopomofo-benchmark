package main

import (
	"bufio"
	"os"
	"strings"
)

func getBenchmarkInput(filename string) (output []BenchmarkInput, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	for scanner := bufio.NewScanner(fd); scanner.Scan(); {
		text := scanner.Text()

		comment := strings.Index(text, "#")
		if comment != -1 {
			text = text[:comment]
		}

		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
	}

	return output, err
}
