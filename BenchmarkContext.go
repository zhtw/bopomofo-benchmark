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
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

type Accuracy struct {
	wordCount    int
	correctCount int
	expectString string
	actualString string
	bopomofo     string
}

type BenchmarkInput struct {
	inputString   string
	inputBopomofo string
}

type BenchmarkItem interface {
	Deinit()
	GetName() string
	EnterBenchmarkInput(input *BenchmarkInput)
	GetAccuracy() []Accuracy
}

type BenchmarkContext struct {
	benchmarkItem []BenchmarkItem
}

func (ctx *BenchmarkContext) AddBenchmarkItem(item BenchmarkItem) {
	ctx.benchmarkItem = append(ctx.benchmarkItem, item)
}

func (ctx *BenchmarkContext) Deinit() {
	for _, item := range ctx.benchmarkItem {
		item.Deinit()
	}

	ctx.benchmarkItem = ctx.benchmarkItem[:0]
}

func (ctx *BenchmarkContext) EnterBenchmarkInput(input *BenchmarkInput) {
	for _, item := range ctx.benchmarkItem {
		item.EnterBenchmarkInput(input)
	}
}

func (ctx *BenchmarkContext) GenerateReport(reportDir string) {
	for _, item := range ctx.benchmarkItem {
		reportName := item.GetName() + ".csv"
		reportName = filepath.Join(reportDir, reportName)

		fd, err := os.OpenFile(reportName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			panic(fmt.Sprintf("Cannot open %s: %s", reportName, err))
		}
		defer fd.Close()

		writer := csv.NewWriter(fd)
		writer.Write([]string{"expect", "actual", "correctCount", "wordCount"})

		totalCorrectCount := 0
		totalWordCount := 0
		for _, accuracy := range item.GetAccuracy() {
			writer.Write([]string{
				accuracy.expectString,
				accuracy.actualString,
				fmt.Sprintf("%d", accuracy.correctCount),
				fmt.Sprintf("%d", accuracy.wordCount)})
			totalCorrectCount += accuracy.correctCount
			totalWordCount += accuracy.wordCount
		}

		writer.Flush()

		fmt.Printf("- %s: %d / %d (%0.2f)\n",
			item.GetName(),
			totalCorrectCount,
			totalWordCount,
			float64(totalCorrectCount)/float64(totalWordCount))
	}
}
