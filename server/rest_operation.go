package server

import (
	"encoding/json"
	"gokv/core"
	"net/http"
)

func resp(w http.ResponseWriter, data interface{}) {
	header := w.Header()
	header.Add("Content-Type", "application/json;charset=utf-8")

	template := HttpResponseTemplate{Code: 200, Msg: "ok", Data: data}
	enc := json.NewEncoder(w)
	enc.Encode(template)
}

type HttpResponseTemplate struct {
	Code int
	Msg  string
	Data interface{}
}

// parse command from client http request and execute them
type RestOp struct {
	sop core.StoreOperation
	cop core.ClusterOperation
}

/// StoreOperation
func (h *RestOp) Set(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var result = "OK"
	defer resp(w, result)

	key := r.Form.Get("k")
	value := r.Form.Get("v")

	h.sop.Set(key, value)
	resp(w, result)
}

func (h *RestOp) Get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("k")

	result := h.sop.Get(key)
	resp(w, result)
}

func (h *RestOp) Exist(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("k")

	result := h.sop.Exist(key)
	resp(w, result)
}

func (h *RestOp) Remove(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("k")

	result := h.sop.Remove(key)
	resp(w, result)
}

/// StoreOperation end

/// ClusterOperation
func (h *RestOp) Join(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newNode := r.Form.Get("newNode")

	result := h.cop.Join(newNode)
	resp(w, result)
}

/// ClusterOperation end
