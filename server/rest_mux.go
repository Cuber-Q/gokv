package server

import (
	"fmt"
	"gokv/core"
	"log"
	"net"
	"net/http"
	"strconv"
)

type RestMux struct {
	opMap    map[string]func(w http.ResponseWriter, r *http.Request)
	restIp   string
	restPort int
}

func newRestMux(restIp string, restPort int, sop core.StoreOperation, cop core.ClusterOperation) *RestMux {
	return &RestMux{
		opMap:    routerInit(sop, cop),
		restIp:   restIp,
		restPort: restPort,
	}
}

// init http handler router
func routerInit(sop core.StoreOperation, cop core.ClusterOperation) map[string]func(w http.ResponseWriter, r *http.Request) {
	op := &RestOp{
		sop: sop,
		cop: cop,
	}

	opMap := make(map[string]func(w http.ResponseWriter, r *http.Request))
	opMap["/set"] = op.Set
	opMap["/get"] = op.Get
	opMap["/exist"] = op.Exist
	opMap["/remove"] = op.Remove
	opMap["/join"] = op.Join

	return opMap
}

// start the rest server
func (m *RestMux) start() {
	l, err := net.Listen("tcp", m.restIp+":"+strconv.Itoa(m.restPort))
	if err != nil {
		log.Fatal(fmt.Sprintf("listen http failed"))
	}

	go func() {
		http.Serve(l, m)
	}()
}

// implementation of http/server.go
// A Handler responds to an HTTP request.
func (m *RestMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if fuc := m.opMap[r.URL.Path]; fuc != nil {
		fuc(w, r)
		return
	}

	http.NotFound(w, r)
	return
}
