package main

import (
	"github.com/valyala/fasthttp"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	// load functions of all routers
	routers = []struct {
		name string
		load func(routes []route) http.Handler
	}{
		{"Ace", loadAce},
		{"Aero", loadAero},
		{"Bear", loadBear},
		{"Beego", loadBeego},
		{"Bone", loadBone},
		{"Chi", loadChi},
		{"CloudyKitRouter", loadCloudyKitRouter},
		{"Denco", loadDenco},
		{"Echo", loadEcho},
		{"Gin", loadGin},
		{"GocraftWeb", loadGocraftWeb},
		{"Goji", loadGoji},
		{"Gojiv2", loadGojiv2},
		{"GoJsonRest", loadGoJsonRest},
		{"GoRestful", loadGoRestful},
		{"GorillaMux", loadGorillaMux},
		{"GowwwRouter", loadGowwwRouter},
		{"HttpRouter", loadHttpRouter},
		{"HttpTreeMux", loadHttpTreeMux},
		//{"Kocha", loadKocha},
		{"LARS", loadLARS},
		{"Macaron", loadMacaron},
		{"Martini", loadMartini},
		{"Pat", loadPat},
		{"Possum", loadPossum},
		{"R2router", loadR2router},
		// {"Revel", loadRevel},
		{"Rivet", loadRivet},
		//{"Tango", loadTango},
		{"TigerTonic", loadTigerTonic},
		{"Traffic", loadTraffic},
		{"Vulcan", loadVulcan},
		// {"Zeus", loadZeus},
	}

	// load functions of all fast Routers
	fastRouters = []struct {
		name string
		load func(routes []route) fasthttp.RequestHandler
	}{
		{"Fiber", loadFiber},
	}

	// all APIs
	apis = []struct {
		name   string
		routes []route
	}{
		{"GitHub", githubAPI},
		{"GPlus", gplusAPI},
		{"Parse", parseAPI},
		{"Static", staticRoutes},
	}
)

func TestRouters(t *testing.T) {
	loadTestHandler = true

	for _, router := range routers {
		req, _ := http.NewRequest("GET", "/", nil)
		u := req.URL
		rq := u.RawQuery

		for _, api := range apis {
			r := router.load(api.routes)

			for _, route := range api.routes {
				w := httptest.NewRecorder()
				req.Method = route.method
				req.RequestURI = route.path
				u.Path = route.path
				u.RawQuery = rq
				r.ServeHTTP(w, req)
				if w.Code != 200 || w.Body.String() != route.path {
					t.Errorf(
						"%s in API %s: %d - %s; expected %s %s\n",
						router.name, api.name, w.Code, w.Body.String(), route.method, route.path,
					)
				}
			}
		}
	}

	loadTestHandler = false
}

func TestFastRouters(t *testing.T) {
	loadTestFastHandler = true
	for _, router := range fastRouters {

		ctx := acquireFastCtx("GET", "/")

		for _, api := range apis {
			handler := router.load(api.routes)

			for _, route := range api.routes {
				ctx.Request.Reset()
				ctx.Response.Reset()
				ctx.Request.Header.SetMethod(route.method)
				ctx.Request.Header.SetRequestURI(route.path)
				handler(ctx)
				code, body := ctx.Response.StatusCode(), ctx.Response.Body()
				if code != 200 || string(body) != route.path {
					t.Errorf(
						"%s in API %s: %d - %s; expected %s %s\n",
						router.name, api.name, code, string(body), route.method, route.path,
					)
				}
			}
		}
	}

	loadTestFastHandler = false
}
