package handler

import (
	"gokv/core"
	"net/http"
	"encoding/json"
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

	result := core.Set(key, value)
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