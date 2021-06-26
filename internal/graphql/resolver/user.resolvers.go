package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gigamono/gigamono-api/internal/crud"
	"github.com/gigamono/gigamono-api/internal/graphql/generated"
	"github.com/gigamono/gigamono-api/internal/graphql/model"
)

func (r *sessionUserResolver) Preferences(ctx context.Context, obj *model.SessionUser) (*model.Preferences, error) {
	return crud.GetSessionUserPreferences(ctx, r.App, &obj.ID)
}

// SessionUser returns generated.SessionUserResolver implementation.
func (r *Resolver) SessionUser() generated.SessionUserResolver { return &sessionUserResolver{r} }

type sessionUserResolver struct{ *Resolver }
