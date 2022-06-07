package main

import (
	"github.com/wasmcloud/actor-tinygo"
	httpserver "github.com/wasmcloud/interfaces/httpserver/tinygo"
)

func main() {
	me := Hello{}
	actor.RegisterHandlers(httpserver.HttpServerHandler(&me), actor.ActorHandler(&me))
}

type Hello struct{}

func (e *Hello) HandleRequest(ctx *actor.Context, req httpserver.HttpRequest) (*httpserver.HttpResponse, error) {
	r := httpserver.HttpResponse{
		StatusCode: 200,
		Header:     make(httpserver.HeaderMap, 0),
		Body:       []byte("Hello Security Team\n"),
	}
	return &r, nil
}

func (e *Hello) HealthRequest(ctx *actor.Context, arg actor.HealthCheckRequest) (*actor.HealthCheckResponse, error) {
	var r actor.HealthCheckResponse
	r.Healthy = true
	return &r, nil
}
