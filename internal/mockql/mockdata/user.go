package mockdata

import (
	"github.com/sageflow/sageapi/internal/mockql/model"
)

// GetSessionUser ...
func GetSessionUser() *model.SessionUser {
	return &model.SessionUser{
		ID:      "4e6c8b75-4a74-42fd-84b3-f319a8f153c6",
		Profile: getSessionUserProfile(),
		Session: getSession(),
	}
}
