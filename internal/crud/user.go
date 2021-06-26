package crud

import (
	"context"
	"fmt"

	"github.com/gigamono/gigamono-api/internal/graphql/model"
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gigamono/gigamono/pkg/database/models/resource"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/services/rest/middleware"
	"github.com/gofrs/uuid"
)

// GetSessionUser get the session user main details.
func GetSessionUser(ctx context.Context, app *inits.App) (*model.SessionUser, error) {
	// TODO: Sec: Validation.
	userID := ctx.Value(middleware.SessionDataKey).(middleware.SessionData).User.ID

	// Get the user from db.
	user := resource.User{Base: models.Base{ID: userID}}
	if err := user.GetByID(&app.DB); err != nil {
		panic(errs.NewSystemError("", "getting user", err))
	}

	return &model.SessionUser{
		ID:      userID.String(),
		Profile: CopyProfile(&user.Profile),
	}, nil
}

// GetSessionUserPreferences gets user's preferences from the database.
func GetSessionUserPreferences(_ context.Context, app *inits.App, userID *string) (*model.Preferences, error) {
	// TODO: Sec: Validation, Permission.
	userUUID, err := uuid.FromString(*userID)
	if err != nil {
		panic(err)
	}

	// Get the preferences from db.
	user := resource.User{Base: models.Base{ID: userUUID}}
	preferences, err := user.GetPreferencesByID(&app.DB)
	if err != nil {
		panic(errs.NewSystemError("", "getting preferences", err))
	}

	return CopyPreferences(preferences), nil
}
