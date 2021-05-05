package server

import (
	"fmt"
	"net"

	"github.com/gigamono/gigamono/pkg/logs"
	"github.com/gigamono/gigamono/pkg/services/grpc"
	"github.com/gigamono/gigamono/pkg/services/proto/generated"

	"github.com/go-playground/validator/v10"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"

	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gin-gonic/gin"
)

// APIServer represents an new REST-based server instance.
type APIServer struct {
	inits.App
	*gin.Engine
	Validate   validator.Validate
	AuthClient generated.AuthClient
}

// NewAPIServer creates a new server instance.
func NewAPIServer(app inits.App) (APIServer, error) {
	validate := *validator.New()
	var authClient generated.AuthClient

	// TODO: Sec: Insecure connection.
	client, err := grpc.GetInsecureClient("localhost", app.Config.Services.Types.Auth.Port, app.Config)
	if err != nil {
		logs.FmtPrintln("initialising API server: unable to connect to Auth Service:", err)
	} else {
		authClient = client.(generated.AuthClient)
	}

	return APIServer{
		App:        app,
		Engine:     gin.Default(),
		Validate:   validate,
		AuthClient: authClient,
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
