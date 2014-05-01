package main

type Accuracy struct {
	wordCount    int
	correctCount int
}

type BenchmarkItem interface {
	deinit()
	getName() string
	addAccuracy(accuracy Accuracy)
}

type BenchmarkContext_ struct {
	benchmarkItem []BenchmarkItem
}

func (ctx *BenchmarkContext_) addBenchmarkItem(item BenchmarkItem) {
	ctx.benchmarkItem = append(ctx.benchmarkItem, item)
}

func (ctx *BenchmarkContext_) deinit() {
	for _, item := range ctx.benchmarkItem {
		item.deinit()
	}

	ctx.benchmarkItem = ctx.benchmarkItem[:0]
}
