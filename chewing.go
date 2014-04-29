package main

// #cgo CFLAGS: -I/usr/include/chewing
// #cgo LDFLAGS: -lchewing
// #include <chewing.h>
import "C"

type ChewingContext struct {
	BenchmarkContext
	chewingContext *[0]byte
}

func NewChewingContext() (context *ChewingContext) {
	ctx := new(ChewingContext)
	ctx.name = "chewing"
	ctx.chewingContext = C.chewing_new2(nil, nil, nil, nil)
	return ctx
}

func (c *ChewingContext) deleteChewingContext() {
	if c != nil {
		C.chewing_delete(c.chewingContext)
		c.chewingContext = nil
	}
}
