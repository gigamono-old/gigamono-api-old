package main

import (
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/logs"

	"github.com/gigamono/gigamono-api/pkg/server"
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
