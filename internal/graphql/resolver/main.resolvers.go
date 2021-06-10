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

func (r *mutationResolver) CreateIntegration(ctx context.Context, specification string) (*model.Integration, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UploadIntegrationAvatar(ctx context.Context, id string, file graphql.Upload) (*model.Integration, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PatchIntegrationSpecification(ctx context.Context, id string, patch string) (*model.Integration, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateWorkflow(ctx context.Context, specification string) (*model.Workflow, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PatchWorkflowSpecification(ctx context.Context, id string, patch string) (*model.Workflow, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetSessionUser(ctx context.Context) (*model.SessionUser, error) {
	return crud.GetSessionUser(ctx)
}

func (r *queryResolver) GetIntegration(ctx context.Context, id string) (*model.Integration, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
