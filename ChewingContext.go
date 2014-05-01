package main

import (
	"fmt"
)

// #cgo CFLAGS: -I/usr/include/chewing
// #cgo LDFLAGS: -lchewing
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

func newChewingBenchmarkItem() *ChewingBenchmarkContext {
	ctx := new(ChewingBenchmarkContext)

	ctx.ctx = C.chewing_new2(nil, nil, nil, nil)
	if ctx.ctx == nil {
		panic("chewing_new2 returns NULL")
	}

	return ctx
}

func (ctx *ChewingBenchmarkContext) deinit() {
	if ctx.ctx == nil {
		return
	}

	C.chewing_delete(ctx.ctx)
	ctx.ctx = nil
}

func (ctx *ChewingBenchmarkContext) getName() string {
	return "chewing"
}

func (ctx *ChewingBenchmarkContext) enterBenchmarkInput(input *BenchmarkInput) {
	if ctx.ctx == nil {
		return
	}

	ctx.enterBopomofo(input)
	ctx.computeAccuracy(input)
}

func (ctx *ChewingBenchmarkContext) enterBopomofo(input *BenchmarkInput) {
	C.chewing_clean_bopomofo_buf(ctx.ctx)
	C.chewing_clean_preedit_buf(ctx.ctx)

	for _, key := range bopomofoToKey(input.inputBopomofo) {
		C.chewing_handle_Default(ctx.ctx, C.int(key))
	}
}

func (ctx *ChewingBenchmarkContext) computeAccuracy(input *BenchmarkInput) {
	var accuracy Accuracy

	result := C.GoString(C.chewing_buffer_String_static(ctx.ctx))

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

func (ctx *ChewingBenchmarkContext) getAccuracy() []Accuracy {
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
