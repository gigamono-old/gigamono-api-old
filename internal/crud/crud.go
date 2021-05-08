package crud

import (
	"context"

	"github.com/gigamono/gigamono-api/internal/graphql/model"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gigamono/gigamono/pkg/services/auth"
)

// GetSessionUser get the session user main details.
func GetSessionUser(_ context.Context, tokens model.TokensInput) (*model.SessionUser, error) {
	// TODO: Sec: Validation, Auth, Permission.
	id, err := auth.GetSessionUserID(auth.Tokens(tokens))
	if err != nil {
		panic(errs.NewSystemError(
			messages.Error["user-auth"].(string),
			"authenticating user",
			err,
		))
	}

	return &model.SessionUser{
		ID:     id.String(),
		Tokens: &model.Tokens{},
	}, nil
}
