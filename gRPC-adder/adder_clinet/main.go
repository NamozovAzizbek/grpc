package main

import (
	"context"
	"log"

	"github.com/NamozovAzizbek/grpc/add"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("did not connect :%v", err)
	}

	defer conn.Close()

	c := add.NewAdderClient(conn)
	respons, err := c.Add(context.Background(), &add.Request{X: 2, Y: 3})

	if err != nil {
		log.Fatalf("could not add %v", err)
	}
	log.Printf("adding %v",respons.Result)
}
