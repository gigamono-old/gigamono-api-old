package resolver

import (
	"github.com/sageflow/sageflow/pkg/inits"
	"github.com/sageflow/sageflow/pkg/services/proto"
	"github.com/sageflow/sageflow/pkg/validations"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver for tracking query states and holding dependencies like
type Resolver struct {
	*inits.App
	AuthService   proto.AuthServiceClient
	EngineService proto.EngineServiceClient
	State         struct {
		ValidationErrorMessages []validations.Error
	}
}
