package main

import (
	"github.com/valyala/fasthttp"
	"testing"
)

var (
	parseFiber fasthttp.RequestHandler
)

func init() {
	println("#ParseAPI fast Routes:", len(parseAPI))

	calcMem("Fiber", func() {
		parseFiber = loadFiber(parseAPI)
	})

	println()
}

func BenchmarkFiber_ParseStatic(b *testing.B) {
	benchFastRequest(b, parseFiber, acquireFastCtx("GET", "/1/users"))
}

func BenchmarkFiber_ParseParam(b *testing.B) {
	benchFastRequest(b, parseFiber, acquireFastCtx("GET", "/1/classes/go"))
}

func BenchmarkFiber_Parse2Params(b *testing.B) {
	benchFastRequest(b, parseFiber, acquireFastCtx("GET", "/1/classes/go/123456789"))
}

func BenchmarkFiber_ParseAll(b *testing.B) {
	benchFastRoutes(b, parseFiber, parseAPI)
}
