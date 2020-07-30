package main

import (
	"github.com/valyala/fasthttp"
	"sync"
	"testing"
)

func benchFastRequest(b *testing.B, handler fasthttp.RequestHandler, ctx *fasthttp.RequestCtx) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		handler(ctx)
	}
	releaseFastCtx(ctx)
}

func benchFastRoutes(b *testing.B, handler fasthttp.RequestHandler, routes []route) {
	ctx := acquireFastCtx("GET", "/")
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, route := range routes {
			ctx.Request.Reset()
			ctx.Response.Reset()
			ctx.Request.Header.SetMethod(route.method)
			ctx.Request.Header.SetRequestURI(route.path)
			handler(ctx)
		}
	}
	releaseFastCtx(ctx)
}

var fastCtxPool = sync.Pool{New: func() interface{} {
	return &fasthttp.RequestCtx{}
}}

func acquireFastCtx(method, url string) *fasthttp.RequestCtx {
	ctx := fastCtxPool.Get().(*fasthttp.RequestCtx)
	ctx.Request.Header.SetMethod(method)
	ctx.Request.Header.SetRequestURI(url)

	return ctx
}

func releaseFastCtx(ctx *fasthttp.RequestCtx) {
	ctx.Request.Reset()
	ctx.Response.Reset()
	fastCtxPool.Put(ctx)
}

func BenchmarkFiber_Param(b *testing.B) {
	handler := loadFiberSingle("GET", "/user/:name", fiberHandle)

	benchFastRequest(b, handler, acquireFastCtx("GET", "/user/gordon"))
}

func BenchmarkFiber_Param5(b *testing.B) {
	handler := loadFiberSingle("GET", fiveColon, fiberHandle)

	benchFastRequest(b, handler, acquireFastCtx("GET", fiveRoute))
}

func BenchmarkFiber_Param20(b *testing.B) {
	handler := loadFiberSingle("GET", twentyColon, fiberHandle)

	benchFastRequest(b, handler, acquireFastCtx("GET", twentyRoute))
}

func BenchmarkFiber_ParamWrite(b *testing.B) {
	handler := loadFiberSingle("GET", "/user/:name", fiberHandleWrite)

	benchFastRequest(b, handler, acquireFastCtx("GET", "/user/gordon"))
}
