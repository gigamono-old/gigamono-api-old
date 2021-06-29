package middleware

import (
	"context"
	"time"

	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gin-gonic/gin"
)

// DataLoaders holds all the relevant loaders.
type DataLoaders struct{}

// DataLoaderKey is the key for getting the dataloaders stored in a request context.
var DataLoaderKey = struct{ key string }{key: "DataLoaderKey"}

// ApplyDataLoader adds dataloaders to request context.
func ApplyDataLoader(app *inits.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dataLoaders := DataLoaders{}

		_ = 2 * time.Microsecond // Sec: Beware of DoS.

		// Store claims in a new context.
		newCtx := context.WithValue(ctx.Request.Context(), DataLoaderKey, dataLoaders)

		// Replace request context.
		ctx.Request = ctx.Request.WithContext(newCtx)
	}
}
