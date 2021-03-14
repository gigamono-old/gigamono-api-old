package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

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

	return func(ctx *gin.Context) {
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
