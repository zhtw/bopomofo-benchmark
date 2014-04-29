package main

import (
	"flag"
	"fmt"
	"os"
)

type MainContext struct {
	hasChewing     bool
	chewingContext *ChewingContext
}

func initMainContext(mainContext *MainContext) {
	InitChewingContext(mainContext)
}

func deinitMainContext(mainContext *MainContext) {
	DeinitChewingContext(mainContext)
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprint(os.Stderr, r)
		}
	}()

	var mainContext MainContext

	flag.BoolVar(&mainContext.hasChewing, "chewing", false, "Enable libchewing benchmark")

	initMainContext(&mainContext)
	defer deinitMainContext(&mainContext)

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
