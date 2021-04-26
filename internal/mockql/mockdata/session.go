package mockdata

import (
	"github.com/gigamono/gigamono-api/internal/mockql/model"
)

func getSession() *model.Session {
	return &model.Session{
		Layout:                getLayoutPreferences(),
		FocusWorkspaceIndex:   0,
		WorkspaceFocusIndices: getWorkspaceFocusIndices(),
	}
}
