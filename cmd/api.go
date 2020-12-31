package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/sageflow/sageflow-api/internal/graphql"
	"github.com/sageflow/sageflow-api/internal/graphql/generated"
	"github.com/sageflow/sageflow-api/internal/helpers"

	"github.com/sageflow/sageflow-api/internal/database"
)

const defaultPort = ":3000"

func graphqlHandler() gin.HandlerFunc {
	handler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func main() {
	// Set where initialization log output goes.
	helpers.SetStatusLogFile()

	// Load env file.
	helpers.LoadEnvFile()

	// Handle CLI arguments.
	// helpers.ProcessArgs()

	// Set up database.
	db := database.Setup()
	defer db.Close() // Close database on exit,

	// Create router.
	router := gin.Default()

	// Serving files in ../sageflow-ui/dist.
	router.Use(static.Serve("/", static.LocalFile("../sageflow-ui/dist", true)))

	// Setting up GraphQL.
	router.POST("/query", graphqlHandler())

	// Listen on port.
	router.Run(defaultPort)
}
