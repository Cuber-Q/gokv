package config

import (
	"gokv/util"
	"log"
	"strconv"
	"strings"
)

type Context struct {
	Endpoint string
	Ip       string
	Port     int
}

var Ctx = &Context{}

func (ctx *Context) Parse() {
	if len(ctx.Endpoint) == 0 {
		log.Fatal("invalid context.Endpoint")
	}

	endpoint := ctx.Endpoint
	info := strings.Split(endpoint, ":")
	ctx.Ip = info[0]
	ctx.Port, _ = strconv.Atoi(info[1])
	ctx.AssertValid()
}

func (ctx *Context) AssertValid() {
	if !util.ValidIp(ctx.Ip) {
		panic("invalid endpoint, ip: " + ctx.Ip)
	}

	if ctx.Port < 0 || ctx.Port > 65535 {
		panic("invalid endpoint, port: " + strconv.Itoa(ctx.Port))
	}
}
