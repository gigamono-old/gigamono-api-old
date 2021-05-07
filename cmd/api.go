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
		logs.FmtPrintln("initialising api:", err)
		return
	}

	// Start an API gRPC server.
	server, err := server.NewAPIServer(app)
	if err != nil {
		logs.FmtPrintln("initialising server:", err)
		return
	}

	// Listen on specified port.
	if err := server.Listen(); err != nil {
		logs.FmtPrintln("trying to listen on port specified:", err)
	}
}
