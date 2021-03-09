package server

import (
	"errors"
	"fmt"
	"net"

	"github.com/sageflow/sageflow/pkg/configs"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"

	"github.com/gin-gonic/gin"
	"github.com/sageflow/sageapi/internal/proto"
	"github.com/sageflow/sageflow/pkg/inits"
	"google.golang.org/grpc"
)

// APIServer represents an new REST-based server instance.
type APIServer struct {
	*gin.Engine
	inits.App
	AuthServiceClient   proto.AuthServiceClient
	EngineServiceClient proto.EngineServiceClient
}

// NewAPIServer creates a new server instance.
func NewAPIServer(app inits.App) (APIServer, error) {
	authServiceClient, err := getInsecureServiceClient("localhost", 3002, app.Config)
	if err != nil {
		return APIServer{}, err
	}

	engineServiceClient, err := getInsecureServiceClient("localhost", 3001, app.Config)
	if err != nil {
		return APIServer{}, err
	}

	return APIServer{
		Engine:              gin.Default(),
		App:                 app,
		AuthServiceClient:   authServiceClient.(proto.AuthServiceClient),
		EngineServiceClient: engineServiceClient.(proto.EngineServiceClient),
	}, nil
}

// Listen makes the server listen on specified port.
func (server *APIServer) Listen() error {
	// Listener on TCP port.
	listener, err := net.Listen("tcp", fmt.Sprint(":", server.Config.Server.API.Port))
	if err != nil {
		return err
	}

	// Create multiplexer and delegate content-types.
	multiplexer := cmux.New(listener)
	grpcListener := multiplexer.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpListener := multiplexer.Match(cmux.HTTP1Fast())

	// Run servers concurrently and sync errors.
	grp := new(errgroup.Group)
	grp.Go(func() error { return server.grpcServe(grpcListener) })
	grp.Go(func() error { return server.httpServe(httpListener) })
	grp.Go(func() error { return multiplexer.Serve() })
	return grp.Wait()
}

// Sec: The assumption is that the servers will run in the same cluster so HTTPS connnection is not important.
func getInsecureServiceClient(host string, port int, config configs.SageflowConfig) (interface{}, error) {
	conn, err := grpc.Dial(fmt.Sprint(host, ":", port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	switch port {
	case config.Server.Auth.Port:
		return proto.NewAuthServiceClient(conn), nil
	case config.Server.Engine.Port:
		return proto.NewEngineServiceClient(conn), nil
	default:
		return nil, errors.New("While initialising API server: While connecting to other grpc servers: Port is not recognised")
	}
}
