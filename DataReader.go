/*
 * Copyright (c) 2014 ChangZhuo Chen <czchen@gmail.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 */
package main

import (
	"bufio"
	"os"
	"strings"
)

func GetBenchmarkInput(filename string) (output []BenchmarkInput, err error) {
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

		out := strings.Split(text, "|")
		if len(out) != 2 {
			continue
		}

		benchmarkInput := BenchmarkInput{
			out[0],
			out[1],
		}

		output = append(output, benchmarkInput)
	}

	return output, err
}
