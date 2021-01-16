package server

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sageflow/sageapi/internal/graphql"
)

// Server represents an new REST-based server instance.
type Server struct {
	*gin.Engine
	Port string
}

// NewServer creates a new server instance.
func NewServer(port string) Server {
	return Server{
		Engine: gin.Default(),
		Port:   ":" + port,
	}
}

// Listen makes the server listen on specified port.
func (server *Server) Listen(port string) error {
	server.Port = port             // Set port.
	server.setRoutes()             // Set routes.
	return server.Run(server.Port) // Listen on port.
}

func (server *Server) setRoutes() {
	server.Use(static.Serve("/", static.LocalFile("../sageui/dist", true))) // Serving files in ../sageflow-ui/dist.
	server.POST("/query", graphql.Handler())                                     // Setting up GraphQL.
}
