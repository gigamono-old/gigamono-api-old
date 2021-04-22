package mockdata

import "github.com/sageflow/sageapi/internal/mockql/model"

func getSessionProjects() []*model.SessionProject {
	return []*model.SessionProject{
		&model.SessionProject{
			Name:               "Project New",
			ID:                 "62096f2c-8ae6-43ef-8236-d2636c4ae522",
			FocusWorkflowIndex: 0,
			Workflows:          getSessionWorkflows(),
			FocusDocumentIndex: 0,
			Documents:          getSessionDocuments(),
		},
		&model.SessionProject{
			Name:               "Unnamed Project",
			ID:                 "680a4a4e-1b5c-4092-801e-9bf737b33f40",
			FocusWorkflowIndex: 0,
			Workflows:          getSessionWorkflows(),
			FocusDocumentIndex: 0,
			Documents:          getSessionDocuments(),
		},
	}
}
