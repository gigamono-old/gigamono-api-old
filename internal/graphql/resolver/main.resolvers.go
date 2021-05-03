package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gigamono/gigamono-api/internal/crud"
	"github.com/gigamono/gigamono-api/internal/graphql/generated"
	"github.com/gigamono/gigamono-api/internal/graphql/model"
)

func (r *queryResolver) GetSessionUser(ctx context.Context, tokens model.TokensInput) (*model.SessionUser, error) {
	return crud.GetSessionUser(ctx, tokens)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
