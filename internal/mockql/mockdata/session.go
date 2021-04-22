package mockdata

import (
	"github.com/sageflow/sageapi/internal/mockql/model"
)

func getSession() *model.Session {
	return &model.Session{
		LayoutPreferences:   getLayoutPreferences(),
		FocusWorkspaceIndex: 0,
		Workspaces:          getSessionWorkSpaces(),
		Integrations:        getSessionIntegrations(),
	}
}
