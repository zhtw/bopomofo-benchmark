package main

import (
	"os"
)

func readRawData(filename string) (bopomofo []string, err error) {

	fd, err := os.Open(filename)
	if err != nil {
		return bopomofo, err
	}
	defer fd.Close()

	return bopomofo, err
}
