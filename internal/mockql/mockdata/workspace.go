package mockdata

import (
	"github.com/gigamono/gigamono-api/internal/mockql/model"
	"github.com/gigamono/gigamono/pkg/strs"
)

func getWorkspaces() []*model.Workspace {
	return []*model.Workspace{
		&model.Workspace{
			ID:          "25870d97-bc3a-481e-a025-5c822573b822",
			Name:        "First Workspace",
			Avatar32url: strs.GetAddress("avatars/wsp-25870d97-bc3a-481e-a025-5c822573b822/avatar.png"),
			Projects:    getProjects(),
		},
		&model.Workspace{
			ID:       "96af3c98-b374-412b-b6f0-89e9c47a4788",
			Name:     "Second Workspace",
			Projects: getProjects(),
		},
	}
}

func getWorkspaceFocusIndices() []*model.WorkspaceFocus {
	return []*model.WorkspaceFocus{
		&model.WorkspaceFocus{
			FocusProjectIndex:   0,
			ProjectFocusIndices: getProjectFocusIndices(),
		},
		&model.WorkspaceFocus{
			FocusProjectIndex:   0,
			ProjectFocusIndices: getProjectFocusIndices(),
		},
	}
}
