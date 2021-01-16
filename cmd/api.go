package main

import (
	"github.com/sageflow/sagedb/pkg/database"
	"github.com/sageflow/sageutils/pkg/envs"
	"github.com/sageflow/sageutils/pkg/logs"

	"github.com/sageflow/sageapi/internal/server"
)

func main() {
	// Set up log status file and load .env file.
	logs.SetStatusLogFile()
	envs.LoadEnvFile()

	// Connect to database.
	database.Connect()

	// Create and start server.
	serv := server.NewServer("3000") // TODO: Get from .env with a default value.
	serv.Start()
}
