package main

import (
	"github.com/valyala/fasthttp"
	"testing"
)

var (
	githubFiber fasthttp.RequestHandler
)

func init() {
	println("#GithubAPI fast Routes:", len(githubAPI))

	calcMem("Fiber", func() {
		githubFiber = loadFiber(githubAPI)
	})

	println()
}

func BenchmarkFiber_GithubStatic(b *testing.B) {
	benchFastRequest(b, githubFiber, acquireFastCtx("GET", "/user/repos"))
}

func BenchmarkFiber_GithubParam(b *testing.B) {
	benchFastRequest(b, githubFiber, acquireFastCtx("GET", "/repos/julienschmidt/httprouter/stargazers"))
}

func BenchmarkFiber_GithubAll(b *testing.B) {
	benchFastRoutes(b, githubFiber, githubAPI)
}
