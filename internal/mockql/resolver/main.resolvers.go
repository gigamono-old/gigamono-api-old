package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/sageflow/sageapi/internal/mockql/generated"
	"github.com/sageflow/sageapi/internal/mockql/mockdata"
	"github.com/sageflow/sageapi/internal/mockql/model"
)

func (r *queryResolver) GetCurrentUser(ctx context.Context, session *model.SessionTokensInput) (*model.User, error) {
	return mockdata.GetCurrentUser(), nil
}

func (r *queryResolver) GetCurrentUserIntegrations(ctx context.Context, session *model.SessionTokensInput) (*model.UserIntegrations, error) {
	return mockdata.GetCurrentUserIntegrations(), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
