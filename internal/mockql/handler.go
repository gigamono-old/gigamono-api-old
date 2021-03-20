package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sageflow/sageflow/pkg/services/graphql/handlers"

	"github.com/sageflow/sageapi/internal/graphql/generated"
	"github.com/sageflow/sageapi/internal/graphql/resolver"
	"github.com/sageflow/sageflow/pkg/inits"
	"github.com/sageflow/sageflow/pkg/services/graphql/directives"
	"github.com/sageflow/sageflow/pkg/services/graphql/interceptors"
	proto "github.com/sageflow/sageflow/pkg/services/proto/generated"
)

// Handler handles requests to a graphQL route.
func Handler(app *inits.App, validate *validator.Validate, authService proto.AuthServiceClient, engineService proto.EngineServiceClient) gin.HandlerFunc {
	// Initialize handler.
	handler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolver.Resolver{
			App:           app,
			AuthService:   authService,
			EngineService: engineService,
			Validate:      validate,
		},
		Directives: generated.DirectiveRoot{
			Tag: directives.Tag,
		},
	}))

	// Add middlewares.
	handler.Use(interceptors.ErrorModifier{})
	handler.SetRecoverFunc(handlers.PanicHandler)

	return func(ctx *gin.Context) {
		// Sec: Responses escaped by json.Marshal so there is some protection against XSS.
		// - https://github.com/99designs/gqlgen/blob/3a31a752df764738b1f6e99408df3b169d514784/graphql/handler/server.go#L101
		// - https://github.com/99designs/gqlgen/blob/3a31a752df764738b1f6e99408df3b169d514784/graphql/handler/transport/util.go#L13
		// - https://github.com/golang/go/blob/a7526bbf72dfe0fde00d88f85fd6cddccdb3a345/src/encoding/json/encode.go#L194
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// PlaygroundHandler handles requests to a graphql route.
func PlaygroundHandler() gin.HandlerFunc {
	handler := playground.Handler("GraphQL", "/graphql") // Docs can be found at the /query
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
