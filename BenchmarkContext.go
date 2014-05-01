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
	"fmt"
)

type Accuracy struct {
	wordCount    int
	correctCount int
}

type BenchmarkInput struct {
	inputString   string
	inputBopomofo string
}

type BenchmarkItem interface {
	deinit()
	getName() string
	enterBenchmarkInput(input *BenchmarkInput)
	getAccuracy() []Accuracy
}

type BenchmarkContext struct {
	benchmarkItem []BenchmarkItem
}

func (ctx *BenchmarkContext) addBenchmarkItem(item BenchmarkItem) {
	ctx.benchmarkItem = append(ctx.benchmarkItem, item)
}

func (ctx *BenchmarkContext) deinit() {
	for _, item := range ctx.benchmarkItem {
		item.deinit()
	}

	ctx.benchmarkItem = ctx.benchmarkItem[:0]
}

func (ctx *BenchmarkContext) enterBenchmarkInput(input *BenchmarkInput) {
	for _, item := range ctx.benchmarkItem {
		item.enterBenchmarkInput(input)
	}
}

func (ctx *BenchmarkContext) print() {
	for _, item := range ctx.benchmarkItem {
		fmt.Printf("name: %s\n", item.getName())
		for _, accuracy := range item.getAccuracy() {
			fmt.Printf("\t%d / %d\n", accuracy.correctCount, accuracy.wordCount)
		}
	}
}
