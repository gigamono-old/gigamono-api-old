package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"os"

	"github.com/sageflow/sageflow-backend/internal/database"
	"github.com/sageflow/sageflow-backend/internal/graphql/generated"
	"github.com/sageflow/sageflow-backend/internal/graphql/model"
	"github.com/sageflow/sageflow-backend/internal/utils"
)

// Function is called when file is imported
func initializeDatabase() *database.DB {
	// Load env file.
	utils.LoadEnvFile()

	// Setup database.
	mongodbURI := os.Getenv("MONGODB_URI")
	return database.Connect(mongodbURI)
}

var db = initializeDatabase()

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	return db.FindUser(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
