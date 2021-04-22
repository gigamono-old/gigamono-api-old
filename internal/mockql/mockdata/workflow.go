package mockdata

import "github.com/sageflow/sageapi/internal/mockql/model"

func getSessionWorkflows() []*model.SessionWorkflow  {
	return []*model.SessionWorkflow {
		&model.SessionWorkflow {
			Name: "Workflow 001",
			ID:   "b09011c2-0743-4531-8f22-baf00bf1937d",
		},
		&model.SessionWorkflow {
			Name: "Workflow 002",
			ID:   "56f8ef16-aebf-4806-97a9-50592db572db",
		},
	}
}
