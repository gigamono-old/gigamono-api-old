package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/sageflow/sageapi/internal/graphql"
	"github.com/sageflow/sagedb"
	"github.com/sageflow/sageutils"
)

func main() {
	// Set up log status file and load .env file.
	sageutils.SetStatusLogFile()
	sageutils.LoadEnvFile()

	// Connect to database.
	sagedb.Connect()

	// TODO: Get from .env with a default value.
	const port = "3000"

	// Set up router.
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("../sageflow-ui/dist", true))) // Serving files in ../sageflow-ui/dist.
	router.POST("/query", graphql.Handler())  // Setting up GraphQL.
	router.Run(":"+port) // Listen on port.
}
