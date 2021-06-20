package crud

import (
	"context"

	"github.com/gigamono/gigamono-api/internal/graphql/model"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/services/crud"
)

// CreateWorkflow creates a new workflow in the database.
func CreateWorkflow(ctx context.Context, app *inits.App, specification string, automationID string) (*model.Workflow, error) {
	workflow, err := crud.CreateWorkflow(ctx, app, specification, automationID)
	if err != nil {
		return nil, err
	}

	return &model.Workflow{
		ID:                   workflow.ID.String(),
		Name:                 workflow.Name,
		SpecificationFileURL: workflow.SpecificationFileURL,
		CreatorID:            workflow.CreatorID.String(),
		AutomationID:         workflow.AutomationID.String(),
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
		SpecificationFileURL: workflow.SpecificationFileURL,
		CreatorID:            workflow.CreatorID.String(),
		AutomationID:         workflow.AutomationID.String(),
	}, nil
}
