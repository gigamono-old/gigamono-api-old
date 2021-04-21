package mockdata

import "github.com/sageflow/sageapi/internal/mockql/model"

func getCurrentUserProjectsBasic() []*model.Project {
	return []*model.Project{
		&model.Project{
			Name:                  "Project New",
			ID:                    "62096f2c-8ae6-43ef-8236-d2636c4ae522",
			Workflows:             getCurrentUserCurrentProjectWorkflowsBasic(),
			SelectedWorkflowIndex: 0,
		},
		&model.Project{
			Name:                  "Unused Project",
			ID:                    "680a4a4e-1b5c-4092-801e-9bf737b33f40",
			Workflows:             []*model.Workflow{},
			SelectedWorkflowIndex: -1,
		},
	}
}

func getCurrentUserCurrentProjectWorkflowsBasic() []*model.Workflow {
	return []*model.Workflow{
		&model.Workflow{
			Name: "Workflow 001",
			ID:   "b09011c2-0743-4531-8f22-baf00bf1937d",
		},
		&model.Workflow{
			Name: "Workflow 002",
			ID:   "56f8ef16-aebf-4806-97a9-50592db572db",
		},
	}
}
