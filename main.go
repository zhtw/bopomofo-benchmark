package main

import (
	"flag"
	"fmt"
	"os"
)

type Context struct {
	hasChewing     bool
	chewingContext *ChewingContext
}

func setup(ctx *Context) {
	if ctx.hasChewing {
		ctx.chewingContext = NewChewingContext()
	}
}

func cleanup(ctx *Context) {
	if ctx.hasChewing {
		ctx.chewingContext.deleteChewingContext()
	}
}

func main() {
	var ctx Context

	flag.BoolVar(&ctx.hasChewing, "chewing", false, "Enable libchewing benchmark")

	setup(&ctx)
	defer cleanup(&ctx)

	flag.Parse()

	for _, input := range flag.Args() {
		fmt.Printf("Processing %s ... ", input)

		_, err := getBenchmarkInput(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open %s\n", input)
			continue
		}

		fmt.Printf("Done\n")
	}
}
