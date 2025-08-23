package userClient

import (
	"context"
	gen_user "github.com/danilkompaniets/go-chat-common/gen/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func New(addr string) (gen_user.UserServiceClient, *grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	if err != nil {
		log.Printf("Failed to connect to %s: %v", addr, err)
		return nil, nil, err
	}

	c := gen_user.NewUserServiceClient(conn)
	return c, conn, nil
}
