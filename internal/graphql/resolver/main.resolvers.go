package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gigamono/gigamono-api/internal/crud"
	"github.com/gigamono/gigamono-api/internal/graphql/generated"
	"github.com/gigamono/gigamono-api/internal/graphql/model"
)

func (r *mutationResolver) UpdateProfile(ctx context.Context, id string, profile model.ProfileInput) (*model.Profile, error) {
	return crud.UpdateProfile(ctx, r.App, &id, &profile)
}

func (r *mutationResolver) UploadProfileAvatar(ctx context.Context, id string, file graphql.Upload) (*string, error) {
	return crud.UploadProfileAvatar(ctx, r.App, &id, &file)
}

func (r *mutationResolver) CreateIntegration(ctx context.Context, integration model.IntegrationInput) (*model.Integration, error) {
	return crud.CreateIntegration(ctx, r.App, &integration)
}

func (r *mutationResolver) UploadIntegrationAvatar(ctx context.Context, id string, file graphql.Upload) (*string, error) {
	return crud.UploadIntegrationAvatar(ctx, r.App, &id, &file)
}

func (r *mutationResolver) PatchIntegrationSpecification(ctx context.Context, id string, patch string) (*model.Integration, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SessionUser(ctx context.Context) (*model.SessionUser, error) {
	return crud.GetSessionUser(ctx, r.App)
}

func (r *queryResolver) Integration(ctx context.Context, id string) (*model.Integration, error) {
	return crud.GetIntegration(ctx, r.App, &id)
}

func (r *queryResolver) Preferences(ctx context.Context, id string) (*model.Preferences, error) {
	return crud.GetPreferences(ctx, r.App, &id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
