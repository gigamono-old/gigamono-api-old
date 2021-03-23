package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sageflow/sageapi/internal/mockql/generated"
	"github.com/sageflow/sageapi/internal/mockql/mockdata"
	"github.com/sageflow/sageapi/internal/mockql/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user model.NewUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	return []*model.User{}, nil
}

func (r *queryResolver) GetCurrentUser(ctx context.Context) (*model.User, error) {
	return mockdata.GetCurrentUser(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
