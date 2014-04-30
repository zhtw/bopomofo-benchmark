package main

import (
	"fmt"
)

// #cgo CFLAGS: -I/usr/include/chewing
// #cgo LDFLAGS: -lchewing
// #include <chewing.h>
import "C"

type ChewingContext struct {
	BenchmarkContext
	ctx *[0]byte
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

func (mainCtx *MainContext) initChewingContext() {
	if !mainCtx.hasChewing {
		return
	}

	mainCtx.chewingContext = new(ChewingContext)
	mainCtx.chewingContext.name = "chewing"

	mainCtx.chewingContext.ctx = C.chewing_new2(nil, nil, nil, nil)
	if mainCtx.chewingContext.ctx == nil {
		panic("chewing_new2 returns NULL")
	}
}

func (mainCtx *MainContext) deinitChewingContext() {
	if mainCtx.chewingContext == nil {
		return
	}

	C.chewing_delete(mainCtx.chewingContext.ctx)
	mainCtx.chewingContext.ctx = nil
}

func (mainCtx *MainContext) enterBenchmarkInput(input *BenchmarkInput) {
	if !mainCtx.hasChewing {
		return
	}

	if mainCtx.chewingContext == nil || mainCtx.chewingContext.ctx == nil {
		panic("mainCtx.chewingContext == nil || mainCtx.chewingContext.ctx == nil")
	}

	_ = bopomofoToKey(input.inputBopomofo)
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
