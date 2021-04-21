package mockdata

import (
	"github.com/sageflow/sageapi/internal/mockql/model"
	"github.com/sageflow/sageflow/pkg/strs"
)

// getCurrentUserWorkspaces ...
func getCurrentUserWorkspaces() []*model.Workspace {
	return []*model.Workspace{
		&model.Workspace{
			Name:        "First Workspace",
			Avatar32url: strs.GetAddress("avatars/wsp-25870d97-bc3a-481e-a025-5c822573b822/avatar.png"),
		},
		&model.Workspace{
			Name:        "Second Workspace",
		},
	}
}
