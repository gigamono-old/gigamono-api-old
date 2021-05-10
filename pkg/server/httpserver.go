package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gigamono/gigamono-api/internal/graphql"
	"github.com/gin-gonic/contrib/static"
)

func (server *APIServer) httpServe() error {
	listener, err := net.Listen(
		"tcp",
		fmt.Sprint(":", server.Config.Services.Types.API.PublicPort),
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
	graphqlHandler := graphql.Handler(&server.App, &server.Validate, server.AuthClient)
	playgroundHandler := graphql.PlaygroundHandler()

	server.Use(static.Serve("/", static.LocalFile("../sageui/dist", true)))   // Serves files.
	server.Use(static.Serve("/avatars", static.LocalFile("./avatars", true))) // Serves avatar images.
	server.POST("/graphql", graphqlHandler)                                   // Handles all graphql requests.
	server.GET("/graphql", graphqlHandler)                                    // Handles query-only graphql requests.
	server.GET("/playground", playgroundHandler)                              // Shows playground UI.
}
