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
	"path/filepath"
	"runtime"
	"testing"
)

func Test_GetBenchmarkInput(t *testing.T) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("Cannot get current filename")
	}

	benchmarkInput := filepath.Join(filepath.Dir(filename), "test", "data", "benchmark-input")

	output, err := GetBenchmarkInput(benchmarkInput)
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
