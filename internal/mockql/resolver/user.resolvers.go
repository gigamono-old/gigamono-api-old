package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sageflow/sageapi/internal/graphql/generated"
	"github.com/sageflow/sageapi/internal/graphql/model"
	gql "github.com/sageflow/sageflow/pkg/services/graphql"
	proto "github.com/sageflow/sageflow/pkg/services/proto/generated"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user model.UserInput) (*model.User, error) {
	// Check for validation errors.
	if err := gql.ValidateStructAndAppendErrors(ctx, r.Validate, user, "user"); err != nil {
		return nil, nil
	}

	// Return error if AuthService is nil.
	if r.AuthService == nil {
		panic(fmt.Errorf("in graphql createUser: AuthServiceClient is not set"))
	}

	// Get token from AuthService/GetSignUpToken.
	response, err := r.AuthService.GetSignInToken(ctx, &proto.UserTokenRequest{})
	if err != nil {
		panic(fmt.Errorf("in graphql createUser: unable to get signin token: %s", err))
	}

	// Send verification email if enabled. Otherwise just create user.
	// r.EmailDriver.Send(...)

	// Might return user or not depending on config.
	return &model.User{
		ID: "xyz",
		Tokens: &model.SessionTokens{
			AccessToken: response.AccessToken,
		},
	}, nil
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
