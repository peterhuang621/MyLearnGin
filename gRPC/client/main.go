package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	pb "gin_demo/gRPC/server/proto"
)

type ClientTokenAuth struct {
	credentials.PerRPCCredentials
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"appID": "peter", "appKey": "12123"}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	// creds, _ := credentials.NewClientTLSFromFile("/Users/peterhuang98/test_code/Go/gin_demo/gRPC/key/test.pem", "*.peterhuang.com")
	conn, err := grpc.NewClient("127.0.0.1:9090", opts...)
	if err != nil {
		log.Fatalf("didn't connect: %v", err)
	}

	client := pb.NewSayHelloClient(conn)
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "peterhuang98"})
	fmt.Println(resp.GetResponseMsg())
}
