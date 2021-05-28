package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gigamono/gigamono-api/internal/graphql"
	"github.com/gigamono/gigamono/pkg/configs"
	"github.com/gigamono/gigamono/pkg/services/rest/middleware"
	"github.com/gin-gonic/contrib/static"
)

func (server *APIServer) httpServe() error {
	listener, err := net.Listen(
		"tcp",
		fmt.Sprint(":", server.Config.Services.API.Ports.Public),
	)
	if err != nil {
		return err
	}

	server.setRoutes() // Set routes.

	// Use http server.
	httpServer := &http.Server{
		Handler: server,
	}

	return httpServer.Serve(listener)
}

func (server *APIServer) setRoutes() {
	// Serve home static files.
	server.Use(static.Serve("/", static.LocalFile("../gigamono-ui/dist", true)))

	// Depending on service config, create a local static folder for serving workflow files.
	if server.Config.Filestore.Workflow.Kind == configs.Local {
		// TODO: Permission middleware.
		// Authenticate session user.
		workflowStaticRoute := server.Group("/workflow", middleware.Authenticate(&server.App))
		workflowStaticRoute.StaticFS("/", http.Dir(server.Config.Filestore.Workflow.Path))
	}

	// GraphQL handlers.
	graphqlHandler := graphql.Handler(&server.App, &server.Validate, server.AuthClient)
	server.POST("/graphql", graphqlHandler) // Handles all graphql requests.
	server.GET("/graphql", graphqlHandler)  // Handles query-only graphql requests.
}
