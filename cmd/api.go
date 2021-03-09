package main

import (
	"github.com/sageflow/sageflow/pkg/inits"
	"github.com/sageflow/sageflow/pkg/logs"

	"github.com/sageflow/sageapi/pkg/server"
)

func main() {
	// Initialises app.
	app, err := inits.NewApp("Resource")
	if err != nil {
		logs.FmtPrintln("Unable to initialize api:", err)
		return
	}

	// Start an API gRPC server.
	server, err := server.NewAPIServer(app)
	if err != nil {
		logs.FmtPrintln("Unable to initialize server:", err)
		return
	}

	// Listen on specified port.
	if err := server.Listen(); err != nil {
		logs.FmtPrintln("Unable to listen on port specified:", err)
	}
}
