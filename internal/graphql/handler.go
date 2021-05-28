package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/gigamono/gigamono-api/internal/graphql/generated"
	"github.com/gigamono/gigamono-api/internal/graphql/resolver"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/services/graphql/directives"
	"github.com/gigamono/gigamono/pkg/services/graphql/middleware"
	proto "github.com/gigamono/gigamono/pkg/services/proto/generated"
)

// Handler handles requests to a graphQL route.
func Handler(
	app *inits.App,
	validate *validator.Validate,
	authClient proto.AuthClient,
) gin.HandlerFunc {
	// Initialize handler.
	handler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolver.Resolver{
			App:        app,
			AuthClient: authClient,
			Validate:   validate,
		},
		Directives: generated.DirectiveRoot{
			Tag: directives.Tag,
		},
	}))

	// Add middlewares.
	handler.Use(middleware.ErrorModifier{})
	handler.SetRecoverFunc(middleware.PanicHandler)

	return func(ctx *gin.Context) {
		// Sec: Responses escaped by json.Marshal so there is some protection against XSS.
		// - https://github.com/99designs/gqlgen/blob/3a31a752df764738b1f6e99408df3b169d514784/graphql/handler/server.go#L101
		// - https://github.com/99designs/gqlgen/blob/3a31a752df764738b1f6e99408df3b169d514784/graphql/handler/transport/util.go#L13
		// - https://github.com/golang/go/blob/a7526bbf72dfe0fde00d88f85fd6cddccdb3a345/src/encoding/json/encode.go#L194
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
