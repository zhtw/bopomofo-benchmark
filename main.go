package main

import (
	"flag"
	"fmt"
	"os"
)

type Context struct {
	hasLibchewing     bool
	libchewingContext BenchmarkContext

	hasLibzhuyin bool
	libzhuyin    BenchmarkContext
}

func main() {
	var ctx Context

	flag.BoolVar(&ctx.hasLibchewing, "libchewing", false, "Enable libchewing benchmark")
	flag.BoolVar(&ctx.hasLibzhuyin, "libzhuyin", false, "Enable libzhuyin benchmark")

	flag.Parse()

	for _, input := range flag.Args() {
		fmt.Printf("Processing %s ... ", input)

		_, err := readRawData(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open %s\n", input)
			break
		}

		fmt.Printf("Done\n")
	}
}
