package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/sageflow/sageflow-backend/internal/graphql"
	"github.com/sageflow/sageflow-backend/internal/graphql/generated"
	"github.com/sageflow/sageflow-backend/internal/utils"
)

const defaultPort = ":3000"

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Load env file.
	utils.LoadEnvFile()

	r := gin.Default()

	// Serving files in ../sageflow-ui/dist.
	r.Use(static.Serve("/", static.LocalFile("../sageflow-ui/dist", true)))

	// Setting up GraphQL.
	r.POST("/query", graphqlHandler())
	r.GET("/play", playgroundHandler())

	// Listen on port.
	r.Run(defaultPort)
}
