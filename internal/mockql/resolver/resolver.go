package resolver

import (
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/services/proto/generated"
	"github.com/go-playground/validator/v10"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver for tracking query states and holding dependencies like App and Service Connections.
type Resolver struct {
	*inits.App
	AuthService           generated.AuthServiceClient
	WorkflowEngineService generated.WorkflowEngineServiceClient
	Validate              *validator.Validate
}
