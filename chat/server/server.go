package main

import (
	"log"
	"net"

	"github.com/NamozovAzizbek/grpc/chat/protocol"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed listen to port :8080 %v", err)
	}

	s := protocol.Server{}
	grpcServer := grpc.NewServer()

	protocol.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Faild to serve gRPC server over port 8080: %v", err)
	}
}
