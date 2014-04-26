package main

type SentenceAccuracy struct {
	wordCount    int
	correctCount int
}

type BenchmarkContext struct {
	name     string
	accuracy []SentenceAccuracy
}

type BenchmarkInput struct {
	inputString   string
	inputBopomofo string
}
