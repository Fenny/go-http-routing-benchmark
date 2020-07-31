package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/gofiber/fiber"
	"github.com/valyala/fasthttp"
)

// fiber
func fiberHandle(_ *fiber.Ctx) {}

func fiberHandleWrite(c *fiber.Ctx) {
	c.Fasthttp.WriteString(c.Params("name"))
}

func fiberHandleTest(c *fiber.Ctx) {
	c.Fasthttp.Write(c.Fasthttp.Request.RequestURI())
}

func loadFiberSingle(method, path string, handle fiber.Handler) fasthttp.RequestHandler {
	router := fiber.New(&fiber.Settings{
		CaseSensitive: true,
		StrictRouting: true,
	})
	router.Add(method, path, handle)
	return router.Handler()
}

func loadFiber(routes []route) fasthttp.RequestHandler {
	h := fiberHandle
	if loadTestFastHandler {
		h = fiberHandleTest
	}

	router := fiber.New(&fiber.Settings{
		CaseSensitive: true,
		StrictRouting: true,
	})
	for _, route := range routes {
		router.Add(route.method, route.path, h)
	}
	return router.Handler()
}

// fasthttprouter
func fasthttprouterHandle(_ *fasthttp.RequestCtx) {}

func fasthttprouterHandleWrite(c *fasthttp.RequestCtx) {
	c.WriteString(c.UserValue("name").(string))
}

func fasthttprouterHandleTest(c *fasthttp.RequestCtx) {
	c.Write(c.Request.RequestURI())
}

func loadFasthttprouterSingle(method, path string, handle fasthttp.RequestHandler) fasthttp.RequestHandler {
	router := fasthttprouter.New()
	router.Handle(method, path, handle)
	return router.Handler
}

func loadFasthttprouter(routes []route) fasthttp.RequestHandler {
	h := fasthttprouterHandle
	if loadTestFastHandler {
		h = fasthttprouterHandleTest
	}

	router := fasthttprouter.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router.Handler
}
