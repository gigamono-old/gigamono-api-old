package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"

	"github.com/sageflow/sageapi/internal/graphql/generated"
)

// Handler handles requests to a graphql route.
func Handler() gin.HandlerFunc {
	handler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{}}))
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
