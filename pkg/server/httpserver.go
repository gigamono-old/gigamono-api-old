package server

import (
	"net"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gigamono/gigamono-api/internal/mockql"
)

func (server *APIServer) httpServe(listener net.Listener) error {
	server.setRoutes() // Set routes.

	// Use http server.
	httpServer := &http.Server{
		Handler: server,
	}

	return httpServer.Serve(listener)
}

func (server *APIServer) setRoutes() {
	graphqlHandler := mockql.Handler(&server.App, &server.Validate, server.AuthServiceClient, server.EngineServiceClient)
	playgroundHandler := mockql.PlaygroundHandler()

	server.Use(static.Serve("/", static.LocalFile("../sageui/dist", true)))   // Serves files.
	server.Use(static.Serve("/avatars", static.LocalFile("./avatars", true))) // Serves avatar images.
	server.POST("/graphql", graphqlHandler)                                   // Handles all graphql requests.
	server.GET("/graphql", graphqlHandler)                                    // Handles query-only graphql requests.
	server.GET("/playground", playgroundHandler)                              // Shows playground UI.
}
