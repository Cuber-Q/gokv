package config

import (
	"testing"
	"log"
)

func TestConfig(t *testing.T) {
	ctx := &Context{
		Endpoint:"127.0.0_1:9901",
	}

	ctx.Parse()

	log.Println("context:\n", ctx)
}
