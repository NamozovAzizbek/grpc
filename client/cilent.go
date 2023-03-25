package main

import (
	"context"
	"log"

	"github.com/NamozovAzizbek/grpc/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// grpc.WithTransportCredentials(insecure.NewCredentials())
// grpc.WithInsecure
func main() {
	var err error
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{Body: "Hello from the client!"}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Errot say calling SayHello %s", err)
	}
	log.Printf("Response From server: %s", response.Body)
}
