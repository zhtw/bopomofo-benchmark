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
	"os"
	"path/filepath"
)

// #cgo pkg-config: chewing
// #include <chewing.h>
import "C"

type ChewingBenchmarkContext struct {
	accuracy []Accuracy
	ctx      *[0]byte
}

var BOPOMOFO_START = map[rune]uint8{
	'ㄅ': '1',
	'ㄆ': 'q',
	'ㄇ': 'a',
	'ㄈ': 'z',
	'ㄉ': '2',
	'ㄊ': 'w',
	'ㄋ': 's',
	'ㄌ': 'x',
	'ㄍ': 'e',
	'ㄎ': 'd',
	'ㄏ': 'c',
	'ㄐ': 'r',
	'ㄑ': 'f',
	'ㄒ': 'v',
	'ㄓ': '5',
	'ㄔ': 't',
	'ㄕ': 'g',
	'ㄖ': 'b',
	'ㄗ': 'y',
	'ㄘ': 'h',
	'ㄙ': 'n',
}

var BOPOMOFO_END = map[rune]uint8{
	'ㄧ': 'u',
	'ㄨ': 'j',
	'ㄩ': 'm',

	'ㄚ': '8',
	'ㄛ': 'i',
	'ㄜ': 'k',
	'ㄝ': ',',
	'ㄞ': '9',
	'ㄟ': 'o',
	'ㄠ': 'l',
	'ㄡ': '.',
	'ㄢ': '0',
	'ㄣ': 'p',
	'ㄤ': ';',
	'ㄥ': '/',
	'ㄦ': '-',
}

var BOPOMOFO_TONE = map[rune]uint8{
	'˙': '7',
	'ˊ': '6',
	'ˇ': '3',
	'ˋ': '4',
}

func NewChewingBenchmarkItem(workDir string) *ChewingBenchmarkContext {
	ctx := new(ChewingBenchmarkContext)

	workDir = filepath.Join(workDir, "chewing")
	err := os.MkdirAll(workDir, 0700)
	if err != nil {
		panic(fmt.Sprintf("Cannot create directory %s: %s", workDir, err))
	}

	workDir = filepath.Join(workDir, "chewing.sqlite3")

	ctx.ctx = C.chewing_new2(nil, C.CString(workDir), nil, nil)
	if ctx.ctx == nil {
		panic("chewing_new2 returns NULL")
	}

	return ctx
}

func (ctx *ChewingBenchmarkContext) Deinit() {
	if ctx.ctx == nil {
		return
	}

	C.chewing_delete(ctx.ctx)
	ctx.ctx = nil
}

func (ctx *ChewingBenchmarkContext) GetName() string {
	return "chewing"
}

func (ctx *ChewingBenchmarkContext) EnterBenchmarkInput(input *BenchmarkInput) {
	if ctx.ctx == nil {
		return
	}

	ctx.enterBopomofo(input)
	ctx.computeAccuracy(input)
	ctx.selectCandidate(input)
}

func (ctx *ChewingBenchmarkContext) enterBopomofo(input *BenchmarkInput) {
	C.chewing_clean_bopomofo_buf(ctx.ctx)
	C.chewing_clean_preedit_buf(ctx.ctx)

	for _, key := range bopomofoToKey(input.inputBopomofo) {
		C.chewing_handle_Default(ctx.ctx, C.int(key))
	}
}

func (ctx *ChewingBenchmarkContext) computeAccuracy(input *BenchmarkInput) {
	var accuracy Accuracy = Accuracy{
		expectString: input.inputString,
		bopomofo:     input.inputBopomofo,
	}

	result := C.GoString(C.chewing_buffer_String_static(ctx.ctx))
	accuracy.actualString = result

	if len(result) != len(input.inputString) {
		panic("len(result) != len(input.inputString)")
	}

	for i := range result {
		if result[i] == input.inputString[i] {
			accuracy.correctCount++
		}
		accuracy.wordCount++
	}

	ctx.accuracy = append(ctx.accuracy, accuracy)
}

func (ctx *ChewingBenchmarkContext) selectCandidate(input *BenchmarkInput) {
	var ret C.int

	ret = C.chewing_handle_Home(ctx.ctx)
	if ret == -1 {
		panic(fmt.Sprintf("C.chewing_handle_Home(ctx.ctx) = %d", ret))
	}

	for _, word := range input.inputString {

		ret = C.chewing_cand_open(ctx.ctx)
		if ret != 0 {
			panic(fmt.Sprintf("C.chewing_cand_open(ctx.ctx) = %d", ret))
		}

		ret = C.chewing_cand_list_last(ctx.ctx)
		if ret != 0 {
			panic(fmt.Sprintf("C.chewing_cand_list_last(ctx.ctx) = %d", ret))
		}

		total := C.chewing_cand_TotalChoice(ctx.ctx)
		for index := C.int(0); index < total; index++ {
			cand := []rune(C.GoString(C.chewing_cand_string_by_index_static(ctx.ctx, index)))

			if len(cand) != 1 {
				panic("C.chewing_cand_string_by_index_static(ctx.ctx, index) does not return single word")
			}

			if cand[0] == word {
				C.chewing_cand_choose_by_index(ctx.ctx, index)
				break
			}
		}

		ret = C.chewing_handle_Right(ctx.ctx)
		if ret != 0 {
			panic(fmt.Sprintf("C.chewing_handle_Right(ctx.ctx) = %d", ret))
		}
	}

	if C.GoString(C.chewing_buffer_String_static(ctx.ctx)) != input.inputString {
		panic("Cannot select correct word")
	}

	ret = C.chewing_commit_preedit_buf(ctx.ctx)
	if ret != 0 {
		panic(fmt.Sprintf("C.chewing_commit_preedit_buf(ctx.ctx) = %d", ret))
	}
}

func (ctx *ChewingBenchmarkContext) GetAccuracy() []Accuracy {
	return ctx.accuracy
}

func bopomofoToKey(bopomofo string) (keySequence []uint8) {
	terminated := true
	var key uint8
	var ok bool

	for _, runeValue := range bopomofo {
		key, ok = BOPOMOFO_START[runeValue]
		if ok {
			if !terminated {
				keySequence = append(keySequence, ' ')
				terminated = true
			}
			keySequence = append(keySequence, key)
			continue
		}

		key, ok = BOPOMOFO_END[runeValue]
		if ok {
			terminated = false
			keySequence = append(keySequence, key)
			continue
		}

		key, ok = BOPOMOFO_TONE[runeValue]
		if ok {
			terminated = true
			keySequence = append(keySequence, key)
			continue
		}

		panic(fmt.Sprintf("Unknown bopomofo: %c", runeValue))
	}

	if !terminated {
		keySequence = append(keySequence, ' ')
	}

	return keySequence
}
