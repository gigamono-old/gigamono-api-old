package resolver

import (
	"github.com/go-playground/validator/v10"
	"github.com/sageflow/sageflow/pkg/inits"
	"github.com/sageflow/sageflow/pkg/services/proto/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver for tracking query states and holding dependencies like App and Service Connections.
type Resolver struct {
	*inits.App
	AuthService   generated.AuthServiceClient
	EngineService generated.EngineServiceClient
	Validate      *validator.Validate
}
