package server

import (
	"fmt"
	"net"

	"github.com/sageflow/sageflow/pkg/logs"
	"github.com/sageflow/sageflow/pkg/services"
	"github.com/sageflow/sageflow/pkg/services/proto/generated"

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
	AuthServiceClient   generated.AuthServiceClient
	EngineServiceClient generated.EngineServiceClient
}

// NewAPIServer creates a new server instance.
func NewAPIServer(app inits.App) (APIServer, error) {
	validate := *validator.New()
	var (
		authServiceClient   generated.AuthServiceClient
		engineServiceClient generated.EngineServiceClient
	)

	client, err := services.GetInsecureServiceClient("localhost", 3002, app.Config)
	if err != nil {
		logs.FmtPrintln("initialising API server: unable to connect to Auth Service:", err)
	} else {
		authServiceClient = client.(generated.AuthServiceClient)
	}

	client, err = services.GetInsecureServiceClient("localhost", 3001, app.Config)
	if err != nil {
		logs.FmtPrintln("initialising API server: unable to connect to Engine Service:", err)
	} else {
		engineServiceClient = client.(generated.EngineServiceClient)
	}

	return APIServer{
		Engine:              gin.Default(),
		App:                 app,
		Validate:            validate,
		AuthServiceClient:   authServiceClient,
		EngineServiceClient: engineServiceClient,
	}, nil
}

// Listen makes the server listen on specified port.
func (server *APIServer) Listen() error {
	// Listener on TCP port.
	listener, err := net.Listen("tcp", fmt.Sprint(":", server.Config.Services.Types.API.Port))
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
