package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gigamono/gigamono-api/internal/graphql"
	"github.com/gigamono/gigamono/pkg/services/rest/middleware"
	"github.com/gigamono/gigamono/pkg/services/rest/routes"
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

	// Set local static routes if specified.
	routes.SetLocalStaticRoutes(server.Engine, &server.App)

	// GraphQL handler.
	graphqlHandler := graphql.Handler(&server.App, &server.Validate, server.AuthClient)
	graphqlRoute := server.Group("/graphql", middleware.AuthenticateCreateUser(&server.App))
	{
		graphqlRoute.POST("/", graphqlHandler) // Handles all graphql requests.
		graphqlRoute.GET("/", graphqlHandler)  // Handles query-only graphql requests.
	}
}
