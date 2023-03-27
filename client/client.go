package main

import (
	"context"
	"log"

	"github.com/NamozovAzizbek/grpc/chat/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var err error
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", &err)
	}
	defer conn.Close()
	c := protocol.NewChatServiceClient(conn)

	message := protocol.Message{Body: "hello from client!"}

	response, err := c.Telegram(context.Background(), &message)
	if err != nil {
		log.Fatalf("Errot say calling SayHello %s", err)
	}
	log.Printf("Response From server: %s", response.Body)
}
