package mockdata

import (
	"github.com/sageflow/sageapi/internal/mockql/model"
	"github.com/sageflow/sageflow/pkg/strs"
)

// GetSessionUser ...
func GetSessionUser() *model.SessionUser {
	return &model.SessionUser{
		ID:                            "4e6c8b75-4a74-42fd-84b3-f319a8f153c6",
		Profile:                       getSessionUserProfile(),
		LayoutPreferences:             getSessionUserPreferencesLayout(),
		Workspaces:                    getSessionUserWorkspaces(),
		WorkspaceIndex:                0,
		WorkspaceProjects:             getWorkspaceProjects(),
		WorkspaceProjectIndex:         0,
		WorkspaceProjectWorkflows:     getWorkspaceProjectWorkflows(),
		WorkspaceProjectWorkflowIndex: 0,
		WorkspaceProjectDocuments:     getWorkspaceProjectDocuments(),
		WorkspaceProjectDocumentIndex: 0,
		WorkspaceIntegrations:         getWorkspaceIntegrations(),
	}
}

func getSessionUserPreferencesLayout() *model.LayoutPreferences {
	return &model.LayoutPreferences{
		ActivityBarMainShortcuts:  getSessionUserActivityBarMainShortcuts(),
		ActivityBarSpaceShortcuts: getSessionUserActivityBarSpaceShortcuts(),
		ActivityBarOtherShortcuts: getSessionUserActivityBarOtherShortcuts(),
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

func getSessionUserActivityBarSpaceShortcuts() []*model.ShortcutButton {
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
