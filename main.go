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

func (mainContext *MainContext) initMainContext() {
	mainContext.initChewingContext()
}

func (mainContext *MainContext) deinitMainContext() {
	mainContext.deinitChewingContext()
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprint(os.Stderr, r)
		}
	}()

	var mainContext MainContext

	flag.BoolVar(&mainContext.hasChewing, "chewing", false, "Enable libchewing benchmark")

	mainContext.initMainContext()
	defer mainContext.deinitMainContext()

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
