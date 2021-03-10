package server

import (
	"net"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/sageflow/sageapi/internal/graphql"
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
	server.Use(static.Serve("/", static.LocalFile("../sageui/dist", true)))                                                          // Serving files in ../sageflow-ui/dist.
	server.POST("/query", graphql.QueryHandler(&server.App, &server.Validate, server.AuthServiceClient, server.EngineServiceClient)) // Handling the query route
	server.GET("/playground", graphql.PlaygroundHandler())                                                                           // Handling playground route
}
