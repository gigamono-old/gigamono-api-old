package server

import (
	"fmt"
	"net"

	"github.com/sageflow/sageflow/pkg/logs"
	"github.com/sageflow/sageflow/pkg/services"
	"github.com/sageflow/sageflow/pkg/services/proto"

	"github.com/go-playground/validator/v10"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"

	"github.com/gin-gonic/gin"
	"github.com/sageflow/sageflow/pkg/inits"
)

// APIServer represents an new REST-based server instance.
type APIServer struct {
	*gin.Engine
	inits.App
	Validate            validator.Validate
	AuthServiceClient   proto.AuthServiceClient
	EngineServiceClient proto.EngineServiceClient
}

// NewAPIServer creates a new server instance.
func NewAPIServer(app inits.App) (APIServer, error) {
	validate := *validator.New()

	authServiceClient, err := services.GetInsecureServiceClient("localhost", 3002, app.Config)
	if err != nil {
		logs.FmtPrintln("While initialising API server: Unable to connect to Auth Service:", err)
	}

	engineServiceClient, err := services.GetInsecureServiceClient("localhost", 3001, app.Config)
	if err != nil {
		logs.FmtPrintln("While initialising API server: Unable to connect to Engine Service:", err)
	}

	return APIServer{
		Engine:              gin.Default(),
		App:                 app,
		Validate:            validate,
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
