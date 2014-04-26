package main

type SentenceAccuracy struct {
	wordCount    int
	correctCount int
}

type BenchmarkContext struct {
	name     string
	accuracy []SentenceAccuracy
}
