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
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func prepareDir(dir string) string {
	var err error

	dir, err = filepath.Abs(dir)
	if err != nil {
		panic(fmt.Sprintf("Cannot covert %s to absoluted path", dir))
	}

	err = os.MkdirAll(dir, 0700)
	if err != nil {
		panic(fmt.Sprintf("Cannot create directory %s", dir))
	}

	return dir
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "%s\n", r)
		}
	}()

	var chewing bool
	var workDir string
	var reportDir string

	var ctx BenchmarkContext
	defer func() {
		ctx.deinit()
	}()

	flag.BoolVar(&chewing, "chewing", true, "Enable libchewing benchmark")
	flag.StringVar(&workDir, "workdir", "work", "Set working directory")
	flag.StringVar(&reportDir, "reportdir", "report", "Set report directory")

	workDir = prepareDir(workDir)
	reportDir = prepareDir(reportDir)

	flag.Parse()

	if chewing {
		ctx.addBenchmarkItem(newChewingBenchmarkItem())
	}

	for _, input := range flag.Args() {
		fmt.Printf("Processing %s ... ", input)

		inputSeq, err := getBenchmarkInput(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open %s\n", input)
			continue
		}

		for _, input := range inputSeq {
			ctx.enterBenchmarkInput(&input)
		}

		fmt.Printf("Done\n")
	}

	ctx.print()
}
