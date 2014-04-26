package main

import (
	"flag"
	"fmt"
)

type SentenceAccuracy struct {
	wordCount    int
	correctCount int
}

type BenchmarkContext struct {
	hasLibchewing      bool
	libchewingAccuracy SentenceAccuracy

	hasLibzhuyin      bool
	libzhuyinAccuracy SentenceAccuracy
}

func main() {
	var ctx BenchmarkContext

	flag.BoolVar(&ctx.hasLibchewing, "libchewing", false, "Enable libchewing benchmark")
	flag.BoolVar(&ctx.hasLibzhuyin, "libzhuyin", false, "Enable libzhuyin benchmark")

	flag.Parse()

	for _, input := range flag.Args() {
		fmt.Printf("Processing %s ... ", input)
		fmt.Printf("done\n")
	}
}
