package server

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/sageflow/sageflow/pkg/services/proto/generated"
)

func (server *APIServer) grpcServe(listener net.Listener) error {
	grpcServer := grpc.NewServer() // Create a gRPC server.

	// Register gRPC service.
	generated.RegisterAPIServiceServer(grpcServer, server)
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
