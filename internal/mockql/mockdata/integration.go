package mockdata

import (
	"github.com/sageflow/sageapi/internal/mockql/model"
	"github.com/sageflow/sageflow/pkg/strs"
)

func getWorkspaceIntegrations() *model.WorkspaceIntegrations {
	return &model.WorkspaceIntegrations{
		Integrations: getAllIntegrations(),
		Builtins:     getBuiltinIntegrations(),
	}
}

func getAllIntegrations() []*model.Integration {
	return []*model.Integration{
		&model.Integration{
			Name:        "Gmail",
			ID:          "175885c5-863c-4f71-be70-a6f2942374f8",
			Description: "Lorem ipsum dolor",
			Avatar32url: strs.GetAddress("avatars/ext-175885c5-863c-4f71-be70-a6f2942374f8/avatar-84-gmail.webp"),
			HexColor:    strs.GetAddress(""),
			Spec:        strs.GetAddress(``),
		},
		&model.Integration{
			Name:        "Slack",
			ID:          "1d6c500e-c061-4e41-8c54-26876f3b1f7d",
			Description: "Lorem ipsum dolor",
			Avatar32url: strs.GetAddress("avatars/ext-1d6c500e-c061-4e41-8c54-26876f3b1f7d/avatar-84-slack.webp"),
			HexColor:    strs.GetAddress(""),
			Spec:        strs.GetAddress(``),
		},
		&model.Integration{
			Name:        "Trello",
			ID:          "2fe9249e-ce89-49cb-b4e9-c27386185d84",
			Description: "Lorem ipsum dolor",
			Avatar32url: strs.GetAddress("avatars/ext-2fe9249e-ce89-49cb-b4e9-c27386185d84/avatar-84-trello.webp"),
			HexColor:    strs.GetAddress(""),
			Spec:        strs.GetAddress(``),
		},
		&model.Integration{
			Name:        "Twitter",
			ID:          "58ef9d3b-ab35-4021-b15f-2e1530dd8e05",
			Description: "Lorem ipsum dolor",
			Avatar32url: strs.GetAddress("avatars/ext-58ef9d3b-ab35-4021-b15f-2e1530dd8e05/avatar-84-twitter.webp"),
			HexColor:    strs.GetAddress(""),
			Spec:        strs.GetAddress(``),
		},
		&model.Integration{
			Name:        "Active Campaign",
			ID:          "76540273-b7a1-4b79-8d13-34ffa806b48e",
			Description: "Lorem ipsum dolor",
			Avatar32url: strs.GetAddress("avatars/ext-76540273-b7a1-4b79-8d13-34ffa806b48e/avatar-84-active-campaign.webp"),
			HexColor:    strs.GetAddress(""),
			Spec:        strs.GetAddress(``),
		},
		&model.Integration{
			Name:        "Google Sheet",
			ID:          "9c338e76-1f12-4bb1-a714-ab702225a6a0",
			Description: "Lorem ipsum dolor",
			Avatar32url: strs.GetAddress("avatars/ext-9c338e76-1f12-4bb1-a714-ab702225a6a0/avatar-84-google-sheet.webp"),
			HexColor:    strs.GetAddress(""),
			Spec:        strs.GetAddress(``),
		},
	}
}

func getBuiltinIntegrations() []*model.Integration {
	return []*model.Integration{
		&model.Integration{
			Name:        "Webhook",
			ID:          "dbc63530-8028-4940-bd81-107553e259f6",
			Description: "Lorem ipsum dolor",
			Avatar32url: strs.GetAddress("avatars/ext-dbc63530-8028-4940-bd81-107553e259f6/avatar-webhook.svg"),
			HexColor:    strs.GetAddress(""),
		},
		&model.Integration{
			Name:        "Email",
			ID:          "a9d7f137-694b-4f9f-a674-fca00a1257d9",
			Description: "Lorem ipsum dolor",
			Avatar32url: strs.GetAddress("avatars/ext-a9d7f137-694b-4f9f-a674-fca00a1257d9/avatar-email.svg"),
			HexColor:    strs.GetAddress(""),
		},
		&model.Integration{
			Name:        "Schedule",
			ID:          "c1af8574-8ccc-4601-9b55-56905967ea29",
			Description: "Lorem ipsum dolor",
			Avatar32url: strs.GetAddress("avatars/ext-c1af8574-8ccc-4601-9b55-56905967ea29/avatar-schedule.svg"),
			HexColor:    strs.GetAddress(""),
		},
	}
}
