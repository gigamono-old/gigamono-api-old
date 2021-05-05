package resolver

import (
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/go-playground/validator/v10"

	proto "github.com/gigamono/gigamono/pkg/services/proto/generated"
)

// Resolver holds dependencies like app.
type Resolver struct {
	*inits.App
	AuthClient           proto.AuthClient
	Validate              *validator.Validate
}
