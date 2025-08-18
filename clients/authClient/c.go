package authClient

import (
	"context"
	"log"
	"time"

	gen_auth "github.com/danilkompaniets/go-chat-common/gen/gen-auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New(addr string) (gen_auth.AuthServiceClient, *grpc.ClientConn, error) {
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

	c := gen_auth.NewAuthServiceClient(conn)
	return c, conn, nil
}
