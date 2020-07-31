package main

import (
	"github.com/valyala/fasthttp"
	"testing"
)

var (
	githubFiber          fasthttp.RequestHandler
	githubFasthttprouter fasthttp.RequestHandler
)

func init() {
	println("#GithubAPI fast Routes:", len(githubAPI))

	calcMem("Fiber", func() {
		githubFiber = loadFiber(githubAPI)
	})

	calcMem("Fasthttprouter", func() {
		githubFasthttprouter = loadFasthttprouter(githubAPI)
	})

	println()
}

func BenchmarkFiber_GithubStatic(b *testing.B) {
	benchFastRequest(b, githubFiber, acquireFastCtx("GET", "/user/repos"))
}
func BenchmarkFasthttprouter_GithubStatic(b *testing.B) {
	benchFastRequest(b, githubFasthttprouter, acquireFastCtx("GET", "/user/repos"))
}

func BenchmarkFiber_GithubParam(b *testing.B) {
	benchFastRequest(b, githubFiber, acquireFastCtx("GET", "/repos/julienschmidt/httprouter/stargazers"))
}
func BenchmarkFasthttprouter_GithubParam(b *testing.B) {
	benchFastRequest(b, githubFasthttprouter, acquireFastCtx("GET", "/repos/julienschmidt/httprouter/stargazers"))
}

func BenchmarkFiber_GithubAll(b *testing.B) {
	benchFastRoutes(b, githubFiber, githubAPI)
}
func BenchmarkFasthttprouter_GithubAll(b *testing.B) {
	benchFastRoutes(b, githubFasthttprouter, githubAPI)
}
