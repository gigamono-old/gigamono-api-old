package server

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/gigamono/gigamono/pkg/services/proto/generated"
)

func (server *APIServer) grpcServe() error {
	listener, err := net.Listen(
		"tcp",
		fmt.Sprint(":", server.Config.Services.Types.API.PrivatePort),
	)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer() // Create a gRPC server.

	// Register gRPC service.
	generated.RegisterAPIServer(grpcServer, server)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener) // Listen for requests.
}

// SayHello says Hello
func (server *APIServer) SayHello(ctx context.Context, msg *generated.Message) (*generated.Message, error) {
	serverMsg := "API replies: " + msg.Content
	fmt.Println(serverMsg)
	response := generated.Message{
		Content: serverMsg,
	}
	return &response, nil
}
