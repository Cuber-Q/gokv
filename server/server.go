package server

import (
	"net/http"
	"strconv"
	//"github.com/goinggo/mapstructure"
	"log"
)

type GoKVServer struct {
	ip     string
	port   int
}

func (server *GoKVServer) start() {
	http.ListenAndServe(server.ip +":"+ strconv.Itoa(server.port), newHttpMux())
	log.Println("server has already on ", server.ip, ":", server.port)
}

func Server(ip string, port int) *GoKVServer {
	server := &GoKVServer{
		ip:ip,
		port:port,
	}
	server.start()
	return server
}