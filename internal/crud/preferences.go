package crud

import (
	"context"

	"github.com/gigamono/gigamono-api/internal/graphql/model"
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gigamono/gigamono/pkg/database/models/resource"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/copier"
)

// GetPreferences gets an existing preferences from the database.
func GetPreferences(_ context.Context, app *inits.App, preferencesID *string) (*model.Preferences, error) {
	// TODO: Sec: Validation, Permission.
	preferencesUUID, err := uuid.FromString(*preferencesID)
	if err != nil {
		panic(err)
	}

	// Get the preferences from db.
	preferences := resource.Preferences{Base: models.Base{ID: preferencesUUID}}
	if err = preferences.GetByID(&app.DB); err != nil {
		panic(errs.NewSystemError("", "getting preferences", err))
	}

	return CopyPreferences(&preferences), nil
}

// CopyPreferences copies preferences from one struct to another.
func CopyPreferences(preferences *resource.Preferences) *model.Preferences {
	newPreferences := &model.Preferences{}

	copier.Copy(newPreferences, preferences) // Ignore error.

	return newPreferences
}
