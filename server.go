package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main(){
	lis, err := net.Listen("tcp", ":9000")
	if err != nil{
		log.Fatalf("Faild to listen on port :9000 %v", err)
	}
	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("Faild to serve gRPC server over port 9000: %v", err)
	}
}
//go get -u github.com/golang/protobuf/protoc-gen