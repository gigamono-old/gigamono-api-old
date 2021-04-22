package mockdata

import (
	"github.com/sageflow/sageapi/internal/mockql/model"
	"github.com/sageflow/sageflow/pkg/strs"
)

func getLayoutPreferences() *model.LayoutPreferences {
	return &model.LayoutPreferences{
		ActivityBarMainShortcuts:      getSessionUserActivityBarMainShortcuts(),
		ActivityBarWorkspaceShortcuts: getSessionUserActivityBarWorkspaceShortcuts(),
		ActivityBarOtherShortcuts:     getSessionUserActivityBarOtherShortcuts(),
	}
}

func getSessionUserActivityBarMainShortcuts() []*model.ShortcutButton {
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

func getSessionUserActivityBarWorkspaceShortcuts() []*model.ShortcutButton {
	return []*model.ShortcutButton{
		&model.ShortcutButton{
			IconName: strs.GetAddress("Marketing"),
		},
	}
}
func getSessionUserActivityBarOtherShortcuts() []*model.ShortcutButton {
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
