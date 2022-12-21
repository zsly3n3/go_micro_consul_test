package server

import (
	"context"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
	"go_micro_test/pb/proto"
	"log"
	"time"
)

type Hello struct{}

func (s *Hello) Say(ctx context.Context, req *proto.SayRequest, rsp *proto.SayResponse) error {
	fmt.Println("request:", req.Name)
	rsp.Message = "Hello " + req.Name
	return nil
}

func StartServer() {
	fmt.Println(`server start...`)

	r := consul.NewRegistry(
		registry.Addrs("192.168.100.4:8500"))
	rpcServer := server.NewServer(
		server.Name("test.service"),
		server.Address("0.0.0.0:8001"),
		server.Registry(r),
	)

	err := proto.RegisterHelloHandler(rpcServer, &Hello{})
	if err != nil {
		log.Fatal("err0:", err)
	}

	service := micro.NewService(
		micro.Server(rpcServer),
	)

	//service.Init()

	// Run server
	if err = service.Run(); err != nil {
		log.Fatal("err1:", err)
	}
}

//var a uint64

func GrpcServer() {
	fmt.Println(`grpc server start...`)
	r := consul.NewRegistry(registry.Addrs("192.168.100.4:8500"))

	//
	//node := new(registry.Node)
	//node.Address = `192.168.100.4:8500`
	//nodes := []*registry.Node{node}
	//name := `grpc_test.service`
	//err := r.Deregister(&registry.Service{Name: name, Nodes: nodes})
	//if err != nil {
	//	fmt.Println(`Deregister_err:`, err.Error())
	//}

	//a = 0
	//regCheckFunc := func(ctx context.Context) error {
	//	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " do register check")
	//	a++
	//	if a < 3 {
	//		return nil
	//	}
	//	return errors.New("this not earth")
	//}

	name := `grpc_test.service`
	mp := make(map[string]string)
	mp[`create_time`] = time.Now().Format("2006-01-02 15:04:05")
	_grpcServer := grpc.NewServer(
		server.Name(name),
		server.Address("0.0.0.0:8001"),
		server.Registry(r),
		server.Metadata(mp),
		//server.RegisterCheck(regCheckFunc),
		//server.RegisterInterval(5*time.Second),
		//server.RegisterTTL(10*time.Second),
	)

	err := proto.RegisterHelloHandler(_grpcServer, &Hello{})
	if err != nil {
		log.Fatal("err0:", err)
	}

	service := micro.NewService(
		micro.Server(_grpcServer),
	)

	//service.Init()

	// Run server
	if err = service.Run(); err != nil {
		log.Fatal("err1:", err)
	}
}
