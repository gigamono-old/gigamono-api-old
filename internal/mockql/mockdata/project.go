package mockdata

import "github.com/sageflow/sageapi/internal/mockql/model"

func getWorkspaceProjects() []*model.Project {
	return []*model.Project{
		&model.Project{
			Name:                  "Project New",
			ID:                    "62096f2c-8ae6-43ef-8236-d2636c4ae522",
		},
		&model.Project{
			Name:                  "Unused Project",
			ID:                    "680a4a4e-1b5c-4092-801e-9bf737b33f40",
		},
	}
}
