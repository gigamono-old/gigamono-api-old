package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/sageflow/sageapi/internal/graphql"
	"github.com/sageflow/sageapi/internal/graphql/generated"
	"github.com/sageflow/sageutils"
	"github.com/sageflow/sagedb"
)

const defaultPort = ":3000"

func graphqlHandler() gin.HandlerFunc {
	handler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func main() {
	sageutils.SetStatusLogFile() // Set where initialization log output goes.

	sageutils.LoadEnvFile() // Load env file.

	// processArgs() // Handle CLI arguments.

	db := sagedb.Connect() // Set up database.

	defer db.Close() // Close database on exit.

	router := gin.Default() // Create router.

	router.Use(static.Serve("/", static.LocalFile("../sageflow-ui/dist", true))) // Serving files in ../sageflow-ui/dist.

	router.POST("/query", graphqlHandler()) // Setting up GraphQL.

	router.Run(defaultPort) // Listen on port.
}
