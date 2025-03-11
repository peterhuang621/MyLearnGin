package main

import (
	"context"
	"errors"
	"fmt"
	pb "gin_demo/gRPC/server/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (c *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no token transfer")
	}
	var appID, appKey string
	if v, ok := md["appid"]; ok {
		appID = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}
	if appID != "peter" || appKey != "12123" {
		return nil, errors.New("token invalid")
	}

	fmt.Println("hello" + req.RequestName)
	return &pb.HelloResponse{ResponseMsg: "hello for you! " + req.RequestName}, nil
}

func main() {
	// creds, _ := credentials.NewServerTLSFromFile("/Users/peterhuang98/test_code/Go/gin_demo/gRPC/key/test.pem", "/Users/peterhuang98/test_code/Go/gin_demo/gRPC/key/test.key")
	listen, _ := net.Listen("tcp", ":9090")
	// grpcServer := grpc.NewServer(grpc.Creds(creds))
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	pb.RegisterSayHelloServer(grpcServer, &server{})

	err := grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}

}
