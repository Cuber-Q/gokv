package server

import (
	"gokv/server/handler"
	"net/http"
)

type HttpMux struct {
	opMap map[string]func(w http.ResponseWriter, r *http.Request)
}

func (p *HttpMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if fuc := p.opMap[r.URL.Path]; fuc != nil {
		fuc(w, r)
		return
	}

	http.NotFound(w, r)
	return
}

func newHttpMux() *HttpMux {
	return &HttpMux{
		opMap: routerInit(),
	}
}

func routerInit() map[string]func(w http.ResponseWriter, r *http.Request) {
	op := &handler.Op{}

	opMap := make(map[string]func(w http.ResponseWriter, r *http.Request))
	opMap["/set"] = op.Set
	opMap["/get"] = op.Get
	opMap["/exist"] = op.Exist
	opMap["/keys"] = op.Keys
	opMap["/echo"] = op.Echo

	return opMap
}
