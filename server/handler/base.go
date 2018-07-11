package handler

import (
	"encoding/json"
	"gokv/server/mode"
	"net/http"
)

// base http operation handler
type BaseOp struct {
}

func resp(w http.ResponseWriter, data interface{}) {
	header := w.Header()
	header.Add("Content-Type", "application/json;charset=utf-8")

	template := model.HttpResponseTemplate{Code : 200, Msg : "ok", Data: data}
	enc := json.NewEncoder(w)
	enc.Encode(template)
}