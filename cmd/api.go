package main

import (
	"github.com/sageflow/sageflow/pkg/inits"
	"github.com/sageflow/sageflow/pkg/logs"

	"github.com/sageflow/sageapi/internal/server"
)

func main() {
	// Initialises app.
	app, err := inits.NewApp("Resource")
	if err != nil {
		logs.FmtPrintln("Unable to connect to database; Database is needed to continue: ", err)
		return
	}

	// Start an API gRPC server.
	server := server.NewAPIServer(app)
	if err := server.Listen(); err != nil {
		logs.FmtPrintln("Unable to listen on port specified:", err)
	}
}
