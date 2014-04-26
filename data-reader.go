package main

import (
	"os"
)

func getBenchmarkInput(filename string) (output []BenchmarkInput, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	_ = readRawData(filename)

	return output, err
}

func readRawData(filename string) (output []string) {

	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	return output
}
