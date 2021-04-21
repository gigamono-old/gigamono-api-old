package mockdata

import (
	"github.com/sageflow/sageapi/internal/mockql/model"
	"github.com/sageflow/sageflow/pkg/strs"
)

// GetCurrentUser ...
func GetCurrentUser() *model.User {
	return &model.User{
		ID:                     "4e6c8b75-4a74-42fd-84b3-f319a8f153c6",
		Profile:                getCurrentUserProfile(),
		LayoutPreferences:      getCurrentUserPreferencesLayout(),
		Workspaces:             getCurrentUserWorkspaces(),
		SelectedWorkspaceIndex: 0,
		Projects:               getCurrentUserProjectsBasic(),
		SelectedProjectIndex:   0,
	}
}

func getCurrentUserPreferencesLayout() *model.LayoutPreferences {
	return &model.LayoutPreferences{
		ActivityBarMainShortcuts:  getCurrentUserActivityBarMainShortcuts(),
		ActivityBarSpaceShortcuts: getCurrentUserActivityBarSpaceShortcuts(),
		ActivityBarOtherShortcuts: getCurrentUserActivityBarOtherShortcuts(),
	}
}

func getCurrentUserActivityBarMainShortcuts() []*model.ShortcutButton {
	return []*model.ShortcutButton{
		&model.ShortcutButton{
			IconName: strs.GetAddress("dashboard"),
			Route:    strs.GetAddress("dashboard"),
		},
		&model.ShortcutButton{
			IconName: strs.GetAddress("design"),
			Route:    strs.GetAddress("design"),
		},
		&model.ShortcutButton{
			IconName: strs.GetAddress("workflows"),
			Route:    strs.GetAddress("workflows"),
		},
		&model.ShortcutButton{
			IconName: strs.GetAddress("documents-2"),
			Route:    strs.GetAddress("documents"),
		},
		&model.ShortcutButton{
			IconName: strs.GetAddress("extensions-2"),
			Route:    strs.GetAddress("extensions"),
		},
	}
}

func getCurrentUserActivityBarSpaceShortcuts() []*model.ShortcutButton {
	return []*model.ShortcutButton{
		&model.ShortcutButton{
			IconName: strs.GetAddress("Marketing"),
		},
	}
}
func getCurrentUserActivityBarOtherShortcuts() []*model.ShortcutButton {
	return []*model.ShortcutButton{
		&model.ShortcutButton{
			IconName: strs.GetAddress("developer"),
			Route:    strs.GetAddress("developer"),
		},
		&model.ShortcutButton{
			IconName: strs.GetAddress("settings"),
			Route:    strs.GetAddress("settings"),
		},
	}
}
