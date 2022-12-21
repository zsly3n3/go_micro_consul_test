package client

import (
	"context"
	"fmt"
	"github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/consul"
	limiter "github.com/go-micro/plugins/v4/wrapper/ratelimiter/uber"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
	"go_micro_test/custom_hystrix"
	"go_micro_test/pb/proto"
)

func Start() {
	fmt.Println(`client start...`)

	r := consul.NewRegistry(
		registry.Addrs("192.168.100.2:8500"))

	_selector := selector.NewSelector(
		selector.SetStrategy(selector.RoundRobin),
		selector.Registry(r),
	)

	service := micro.NewService(
		micro.Client(client.NewClient()),
		micro.Selector(_selector),
	)

	//service.Init()
	_client := proto.NewHelloService("test.service", service.Client())

	rsp, err := _client.Say(context.TODO(), &proto.SayRequest{Name: "zsly3n"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp)
}

func GrpcStart() {
	fmt.Println(`grpc_client start...`)

	r := consul.NewRegistry(
		registry.Addrs("192.168.100.4:8500"))

	name := `grpc_test.service`

	_selector := selector.NewSelector(
		selector.SetStrategy(selector.RoundRobin),
		selector.Registry(r),
	)

	const QPS = 100
	service := micro.NewService(
		micro.Client(grpc.NewClient()),
		micro.Selector(_selector),
		micro.WrapClient(custom_hystrix.NewClientWrapper()),
		micro.WrapHandler(limiter.NewHandlerWrapper(QPS)),
	)

	//service.Init()

	_client := proto.NewHelloService(name, service.Client())

	rsp, err := _client.Say(context.TODO(), &proto.SayRequest{Name: "babc"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp)

}
