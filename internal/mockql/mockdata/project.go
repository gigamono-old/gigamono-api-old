package mockdata

import "github.com/sageflow/sageapi/internal/mockql/model"

func getProjects() []*model.Project {
	return []*model.Project{
		&model.Project{
			Name:      "Project New",
			ID:        "62096f2c-8ae6-43ef-8236-d2636c4ae522",
			Workflows: getWorkflows(),
			Documents: getDocuments(),
		},
		&model.Project{
			Name:      "Unnamed Project",
			ID:        "680a4a4e-1b5c-4092-801e-9bf737b33f40",
			Workflows: getWorkflows(),
			Documents: getDocuments(),
		},
	}
}

func getProjectFocusIndices() []*model.ProjectFocus {
	return []*model.ProjectFocus{
		&model.ProjectFocus{
			FocusDocumentIndex: 0,
			FocusWorkflowIndex: 0,
		},
		&model.ProjectFocus{
			FocusDocumentIndex: 0,
			FocusWorkflowIndex: 0,
		},
	}
}
