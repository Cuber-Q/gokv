package handler

import (
	"net/http"
	"gokv/gokv-cli/config"
	"time"
)

var (
	Client = &http.Client{Timeout: 10 * time.Second}
)

const (
	SET = "set"
	GET = "get"
	EXIST = "exist"
	KEYS = "keys"
)

func buildUrl(cmdConst string) string {
	url := config.Ctx.Endpoint
	return "http://" + url + "/" + cmdConst
}
