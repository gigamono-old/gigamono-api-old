package server

import (
	"net"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sageflow/sageapi/internal/graphql"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
)

// APIServer represents an new REST-based server instance.
type APIServer struct {
	*gin.Engine
	Port string
}

// NewAPIServer creates a new server instance.
func NewAPIServer() APIServer {
	return APIServer{
		Engine: gin.Default(),
	}
}

// Listen makes the server listen on specified port.
func (server *APIServer) Listen(port string) error {
	// Set port.
	server.Port = port

	// Listener on TCP port.
	listener, err := net.Listen("tcp", ":"+server.Port)
	if err != nil {
		return err
	}

	// Create multiplexer and delegate content-types.
	multiplexer := cmux.New(listener)
	grpcListener := multiplexer.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpListener := multiplexer.Match(cmux.HTTP1Fast())

	// Run servers concurrently and sync errors.
	grp := new(errgroup.Group)
	grp.Go(func() error { return server.grpcServe(grpcListener) })
	grp.Go(func() error { return server.httpServe(httpListener) })
	grp.Go(func() error { return multiplexer.Serve() })
	return grp.Wait()
}

func (server *APIServer) setRoutes() {
	server.Use(static.Serve("/", static.LocalFile("../sageui/dist", true))) // Serving files in ../sageflow-ui/dist.
	server.POST("/query", graphql.Handler())                                // Setting up GraphQL.
}
