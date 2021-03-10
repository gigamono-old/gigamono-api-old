package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/sageflow/sageapi/internal/graphql/generated"
	"github.com/sageflow/sageapi/internal/graphql/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user model.UserInput) (*model.User, error) {
	// Return error if AuthService is nil.
	if r.AuthService == nil {
		return nil, errors.New("Unable to create new user because server is not connected to Auth Service")
	}

	// Validate input.
	fmt.Println(">>> errs =", r.State.ValidationErrorMessages)

	// Get token from AuthService/GetSignUpToken.

	// Return result.
	return &model.User{}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return []*model.User{}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
