package main

import (
	"context"
	"log"
	"net"

	"github.com/NamozovAzizbek/grpc/add"
	"google.golang.org/grpc"
)

type server struct {
	add.UnimplementedAdderServer
}

func (s *server) Add(ctx context.Context, req *add.Request) (*add.Response, error){
log.Println("Recived:", req.X, req.Y)
return &add.Response{Result: req.GetX() + req.GetY()}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	s := grpc.NewServer()

	add.RegisterAdderServer(s, &server{})
	log.Printf("serer listening : %v", l.Addr())

	if err := s.Serve(l); err != nil{
		log.Fatalf("failed to server : %v", err)
	}
}
