package chat

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var err error
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(":9000", grpc.WithInsecure)
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := NewChatServiceClient(conn)

	message := Message{Body: "Hello from the client!"}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil{
		log.Fatalf("Errot say calling SayHello %s", err)
	}
	log.Printf("Response From server: %s", response.Body)
}
