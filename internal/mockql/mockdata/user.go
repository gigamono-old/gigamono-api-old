package mockdata

import (
	"github.com/sageflow/sageapi/internal/mockql/model"
	"github.com/sageflow/sageflow/pkg/strs"
)

// GetCurrentUser ...
func GetCurrentUser() *model.User {
	return &model.User{
		ID:                "4e6c8b75-4a74-42fd-84b3-f319a8f153c6",
		LayoutPreferences: getCurrentUserPreferencesLayout(),
	}
}

func getCurrentUserPreferencesLayout() *model.LayoutPreferences {
	return &model.LayoutPreferences{
		ActivityBarMainShortcuts: getCurrentUserActivityBarMainShortcuts(),
		ActivityBarSpaceShortcuts: getCurrentUserActivityBarSpaceShortcuts(),
		ActivityBarOtherShortcuts: getCurrentUserActivityBarOtherShortcuts(),
	}
}

func getCurrentUserActivityBarMainShortcuts() []*model.Button {
	return []*model.Button{
		&model.Button{
			Name: strs.GetAddress("Dashboard"),
		},
		&model.Button{
			Name: strs.GetAddress("Design"),
		},
		&model.Button{
			Name: strs.GetAddress("Workflow"),
		},
		&model.Button{
			Name: strs.GetAddress("Documents"),
		},
		&model.Button{
			Name: strs.GetAddress("Extensions"),
		},
	}
}
func getCurrentUserActivityBarSpaceShortcuts() []*model.Button {
	return []*model.Button{
		&model.Button{
			Name: strs.GetAddress("Marketing"),
		},
	}
}
func getCurrentUserActivityBarOtherShortcuts() []*model.Button {
	return []*model.Button{
		&model.Button{
			Name: strs.GetAddress("Developer"),
		},
		&model.Button{
			Name: strs.GetAddress("Settings"),
		},
	}
}
