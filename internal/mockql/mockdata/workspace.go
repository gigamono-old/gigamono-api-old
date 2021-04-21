package mockdata

import (
	"github.com/sageflow/sageapi/internal/mockql/model"
	"github.com/sageflow/sageflow/pkg/strs"
)

func getSessionUserWorkspaces() []*model.Workspace {
	return []*model.Workspace{
		&model.Workspace{
			ID:          "25870d97-bc3a-481e-a025-5c822573b822",
			Name:        "First Workspace",
			Avatar32url: strs.GetAddress("avatars/wsp-25870d97-bc3a-481e-a025-5c822573b822/avatar.png"),
		},
		&model.Workspace{
			ID:   "96af3c98-b374-412b-b6f0-89e9c47a4788",
			Name: "Second Workspace",
		},
	}
}
