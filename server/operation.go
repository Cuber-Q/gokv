package server

import (
	"gokv/core"
	"net/http"
	"encoding/json"
	"fmt"
	"time"
	"log"
)

// base http operation handler
type BaseOp struct {
}

func resp(w http.ResponseWriter, data interface{}) {
	header := w.Header()
	header.Add("Content-Type", "application/json;charset=utf-8")

	template := HttpResponseTemplate{Code : 200, Msg : "ok", Data: data}
	enc := json.NewEncoder(w)
	enc.Encode(template)
}

type HttpResponseTemplate struct {
	Code int
	Msg string
	Data interface{}
}

// parse command from client http request and execute them
type Op struct {
	BaseOp
}

func (op *Op) Set(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	value := r.Form.Get("value")


	event := logEntryData{Key: key, Value: value}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		log.Printf("json.Marshal failed, err:%v", err)
		fmt.Fprint(w, "internal error\n")
		return
	}

	var result = "fail"
	applyFuture := server.raft.raft.Apply(eventBytes, 5*time.Second)
	if err := applyFuture.Error(); err != nil {
		log.Printf("raft.Apply failed:%v", err)
		fmt.Fprint(w, "internal error\n")
		return
	}else {
		log.Printf("set op OK, raft.Apply OK")
		fmt.Fprintf(w, "ok\n")
		result = core.Set(key, value)
	}
	resp(w, result)
}

func (op *Op) Get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")

	result := core.Get(key)
	resp(w, result)
}

func (op *Op) Exist(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")

	result := core.Exist(key)
	resp(w, result)
}

func (op *Op) Keys(w http.ResponseWriter, r *http.Request) {
	result := core.Keys()
	resp(w, result)
}

func (op *Op) Echo(w http.ResponseWriter, r *http.Request) {
	resp(w, "echo")
}

func (op *Op) AddNode(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	node := r.Form.Get("node")

	result := AddNode(node)
	resp(w, result)
}
