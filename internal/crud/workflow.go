package crud

import (
	"context"

	"github.com/gigamono/gigamono-api/internal/graphql/model"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/services/crud"
)

// CreateWorkflow creates a new workflow in the database.
func CreateWorkflow(ctx context.Context, app *inits.App, specification string) (*model.Workflow, error) {
	workflow, err := crud.CreateWorkflow(ctx, app, specification)
	if err != nil {
		return nil, err
	}

	return &model.Workflow{
		ID:                   workflow.ID.String(),
		Name:                 workflow.Name,
		IsActive:             &workflow.IsActive,
		CreatorID:            workflow.CreatorID.String(),
		SpecificationFileURL: workflow.SpecificationFileURL,
	}, nil
}

// GetWorkflow gets an existing workflow from the database.
func GetWorkflow(ctx context.Context, app *inits.App, workflowID string) (*model.Workflow, error) {
	workflow, err := crud.GetWorkflow(ctx, app, workflowID)
	if err != nil {
		return nil, err
	}

	return &model.Workflow{
		ID:                   workflow.ID.String(),
		Name:                 workflow.Name,
		IsActive:             &workflow.IsActive,
		CreatorID:            workflow.CreatorID.String(),
		SpecificationFileURL: workflow.SpecificationFileURL,
	}, nil
}
