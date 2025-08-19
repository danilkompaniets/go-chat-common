package chatClient

import (
	gen_chat "github.com/danilkompaniets/go-chat-common/gen/gen-chat"
	"google.golang.org/grpc"
	"log"
	"time"
)

func New(addr string) (*gen_chat.ChatServiceClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Printf("Failed to connect to %s: %v", addr, err)
		return nil, err
	}

	defer conn.Close()

	c := gen_chat.NewChatServiceClient(conn)
	return &c, nil
}
