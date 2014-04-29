package main

// #cgo CFLAGS: -I/usr/include/chewing
// #cgo LDFLAGS: -lchewing
// #include <chewing.h>
import "C"

type ChewingContext struct {
	BenchmarkContext
	ctx *[0]byte
}

var BOPOMOFO_MAPPING = map[rune]uint8{
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
