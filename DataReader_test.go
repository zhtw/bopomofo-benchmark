package main

import (
	"path/filepath"
	"runtime"
	"testing"
)

func Test_getBenchmarkInput(t *testing.T) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("Cannot get current filename")
	}

	benchmarkInput := filepath.Join(filepath.Dir(filename), "test", "data", "benchmark-input")

	output, err := getBenchmarkInput(benchmarkInput)
	if err != nil {
		t.Fatal(err)
	}

	if len(output) != 1 {
		t.Fatalf("len(output) = %d shall be %d", len(output), 1)
	}

	expectString := "這是輸入"
	if output[0].inputString != expectString {
		t.Fatalf("inputString `%s' shall be `%s'", output[0].inputString, expectString)
	}

	expectBopomofo := "ㄓㄜˋㄕˋㄕㄨㄖㄨˋ"
	if output[0].inputBopomofo != expectBopomofo {
		t.Fatalf("inputBopomofo `%s' shall be `%s'", output[0].inputBopomofo, expectBopomofo)
	}
}
