package main

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
