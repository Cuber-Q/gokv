package handler

import (
	"gokv/core"
	"net/http"
)

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