package crud

import (
	"context"

	"github.com/gigamono/gigamono-api/internal/graphql/model"
	"github.com/gigamono/gigamono/pkg/services/rest/middleware"
)

// GetSessionUser get the session user main details.
func GetSessionUser(ctx context.Context) (*model.SessionUser, error) {
	// TODO: Sec: Validation.
	userID := ctx.Value(middleware.SessionDataKey).(middleware.SessionData).Claims.Subject

	return &model.SessionUser{
		ID: userID,
	}, nil
}
