package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/sageflow/sageapi/internal/graphql/generated"
	"github.com/sageflow/sageapi/internal/graphql/resolver"
	"github.com/sageflow/sageflow/pkg/inits"
	"github.com/sageflow/sageflow/pkg/services/proto"
	"github.com/sageflow/sageflow/pkg/validations"
)

// QueryHandler handles requests to a query route.
func QueryHandler(app *inits.App, validate *validator.Validate, authService proto.AuthServiceClient, engineService proto.EngineServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Sec: Each request gets a new config, resolver and handler because resolver contains state that can lead to race condition.

		// Initialize resolver.
		newResolver := resolver.Resolver{
			App:           app,
			AuthService:   authService,
			EngineService: engineService,
		}

		// Initialize config.
		config := generated.Config{Resolvers: &newResolver}

		// Set directive implementation.
		config.Directives.Validate = func(ctx context.Context, obj interface{}, next graphql.Resolver, rules *string) (interface{}, error) {
			fieldValue, err := next(ctx)
			if err != nil {
				return fieldValue, err
			}

			err = validate.Var(fieldValue, *rules)
			if err != nil {
				newResolver.State.ValidationErrorMessages = append(newResolver.State.ValidationErrorMessages, validations.Error{
					FieldPath:  graphql.GetPathContext(ctx).Path().String(),
					FieldName:  *graphql.GetPathContext(ctx).Field,
					FieldValue: fieldValue,
					Rules:      *rules,
				})
			}

			return fieldValue, nil
		}

		// Initialize handler.
		handler := handler.NewDefaultServer(generated.NewExecutableSchema(config))
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// PlaygroundHandler handles requests to a graphql route.
func PlaygroundHandler() gin.HandlerFunc {
	handler := playground.Handler("GraphQL", "/query") // Docs can be found at the /query
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
