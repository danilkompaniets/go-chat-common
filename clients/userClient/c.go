package userClient

import (
	gen_user "github.com/danilkompaniets/go-chat/gen/gen-user"
	"google.golang.org/grpc"
	"log"
	"time"
)

func New(addr string) (*gen_user.UserServiceClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Printf("Failed to connect to %s: %v", addr, err)
		return nil, err
	}

	defer conn.Close()

	c := gen_user.NewUserServiceClient(conn)
	return &c, nil
}
