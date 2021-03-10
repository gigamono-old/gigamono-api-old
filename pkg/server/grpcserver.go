package server

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/sageflow/sageflow/pkg/services/proto"
)

func (server *APIServer) grpcServe(listener net.Listener) error {
	grpcServer := grpc.NewServer() // Create a gRPC server.

	// Register gRPC service.
	proto.RegisterAPIServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener) // Listen for requests.
}

// SayHello says Hello
func (server *APIServer) SayHello(ctx context.Context, msg *proto.Message) (*proto.Message, error) {
	serverMsg := "API replies: " + msg.Content
	fmt.Println(serverMsg)
	response := proto.Message{
		Content: serverMsg,
	}
	return &response, nil
}
