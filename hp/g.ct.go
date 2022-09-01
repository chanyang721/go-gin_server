package hp

import (
	"context"
	"log"
	"net/http"
)

type Context struct {
	context.Context
}

func GCtx(r *http.Request) Context {
	if c, ok := r.Context().(Context); !ok {
		log.Panic("Failed to get custom Context from request")
		return c
	} else {
		return c
	}
}
