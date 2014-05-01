package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "%s\n", r)
		}
	}()

	var chewing bool

	var ctx BenchmarkContext
	defer func() {
		ctx.deinit()
	}()

	flag.BoolVar(&chewing, "chewing", true, "Enable libchewing benchmark")
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
