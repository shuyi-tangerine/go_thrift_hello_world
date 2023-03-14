package main

import (
	"context"
	"fmt"
	"github.com/shuyi-tangerine/go_thrift_hello_world/gen-go/base"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/shuyi-tangerine/go_thrift_hello_world/gen-go/tangerine/go/thrift/hello_world"
)

var defaultCtx = context.Background()

func handleClient(client hello_world.HelloWorldSayer) (err error) {
	res, err := client.Say(defaultCtx, &hello_world.SayRequest{Base: &base.RPCRequest{}})
	if err != nil {
		return err
	}
	fmt.Println("res ==>", res.Base.Code, res.Base.Message)
	fmt.Println("say ==>", res.Content)
	return
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool, cfg *thrift.TConfiguration) error {
	var transport thrift.TTransport
	if secure {
		transport = thrift.NewTSSLSocketConf(addr, cfg)
	} else {
		transport = thrift.NewTSocketConf(addr, cfg)
	}
	transport, err := transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return handleClient(hello_world.NewHelloWorldSayerClient(thrift.NewTStandardClient(iprot, oprot)))
}
