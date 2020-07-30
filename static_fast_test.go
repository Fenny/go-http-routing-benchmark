package main

import (
	"github.com/valyala/fasthttp"
	"testing"
)

var (
	staticFiber fasthttp.RequestHandler
)

func init() {
	println("#Static fast Routes:", len(staticRoutes))

	calcMem("Fiber", func() {
		staticFiber = loadFiber(staticRoutes)
	})

	println()
}

func BenchmarkFiber_StaticAll(b *testing.B) {
	benchFastRoutes(b, staticFiber, staticRoutes)
}
