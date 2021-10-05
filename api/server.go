package api

import (
	"fmt"

	"github.com/valyala/fasthttp"
	"github.com/vic020/go-crawler/consts"
)

type HTTPServer struct {
	version string
}

func (s *HTTPServer) Run(addr string) {
	fasthttp.ListenAndServe(addr, s.FastHTTPHandler)

}

func (s *HTTPServer) Close() {

}

func (s *HTTPServer) FastHTTPHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Path: %v, Version: %v", string(ctx.Path()), s.version)
}

var httpServer *HTTPServer

func newHTTPServer() *HTTPServer {
	return &HTTPServer{
		version: consts.ServerHTTPVersion,
	}
}

func NewHTTPServer() *HTTPServer {
	if httpServer == nil {
		httpServer = newHTTPServer()
	}
	return httpServer
}
