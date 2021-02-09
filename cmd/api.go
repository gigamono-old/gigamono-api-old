package main

import (
	"github.com/sageflow/sagedb/pkg/database"
	"github.com/sageflow/sageflow/pkg/envs"
	"github.com/sageflow/sageflow/pkg/logs"

	"github.com/sageflow/sageapi/internal/server"
)

func main() {
	// Set up log status file and load .env file.
	logs.SetStatusLogFile() // TODO. logs.SetStatusLogFile(config.Logging.Status.Filepath)
	envs.LoadEnvFile()      // TODO. Remove!

	// Connect to database.
	database.Connect() // TODO. db := database.Connect(config.db)

	// Create and start server.
	serv := server.NewServer() // TODO. server.NewServer(db, config)
	serv.Listen("3000") // TODO. database.Connect(config.Server.API.Port)
}
