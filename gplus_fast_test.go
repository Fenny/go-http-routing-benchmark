package main

import (
	"github.com/valyala/fasthttp"
	"testing"
)

var (
	gplusFiber fasthttp.RequestHandler
)

func init() {
	println("#GPlusAPI fast Routes:", len(gplusAPI))

	calcMem("Fiber", func() {
		gplusFiber = loadFiber(gplusAPI)
	})

	println()
}

func BenchmarkFiber_GPlusStatic(b *testing.B) {
	benchFastRequest(b, gplusFiber, acquireFastCtx("GET", "/people"))
}

func BenchmarkFiber_GPlusParam(b *testing.B) {
	benchFastRequest(b, gplusFiber, acquireFastCtx("GET", "/people/118051310819094153327"))
}

func BenchmarkFiber_GPlus2Params(b *testing.B) {
	benchFastRequest(b, gplusFiber, acquireFastCtx("GET", "/people/118051310819094153327/activities/123456789"))
}

func BenchmarkFiber_GPlusAll(b *testing.B) {
	benchFastRoutes(b, gplusFiber, gplusAPI)
}
