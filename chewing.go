package main

// #cgo CFLAGS: -I/usr/include/chewing
// #cgo LDFLAGS: -lchewing
// #include <chewing.h>
import "C"

type ChewingContext struct {
	BenchmarkContext
	ctx *[0]byte
}

func InitChewingContext(mainCtx *MainContext) {
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

func DeinitChewingContext(mainCtx *MainContext) {
	if mainCtx.chewingContext == nil {
		return
	}

	C.chewing_delete(mainCtx.chewingContext.ctx)
	mainCtx.chewingContext.ctx = nil
}
