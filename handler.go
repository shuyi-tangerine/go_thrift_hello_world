package main

import (
	"context"
	"github.com/shuyi-tangerine/go_thrift_hello_world/gen-go/base"
	"github.com/shuyi-tangerine/go_thrift_hello_world/gen-go/tangerine/go/thrift/hello_world"
)

type HelloWorldSayer struct {
}

func NewHelloWorldSayer() hello_world.HelloWorldSayer {
	return &HelloWorldSayer{}
}

func (m *HelloWorldSayer) Say(ctx context.Context, req *hello_world.SayRequest) (res *hello_world.SayResponse, _err error) {
	res = &hello_world.SayResponse{
		Content: "Hello World!",
		Base: &base.RPCResponse{
			Code:    0,
			Message: "ok",
		},
	}
	return
}
