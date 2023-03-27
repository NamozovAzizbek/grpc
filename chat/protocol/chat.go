package protocol

import (
	context "context"
	"log"
)

type Server struct{}

func (s *Server) Telegram(ctx context.Context, req *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", req.Body)
	return &Message{Body: "message from server!"}, nil
}
