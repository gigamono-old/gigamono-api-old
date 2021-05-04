package crud

import (
	"context"
	"fmt"

	"github.com/gigamono/gigamono-api/internal/graphql/model"
	"github.com/gigamono/gigamono/pkg/services/auth"
)

// GetSessionUser get the session user main details.
func GetSessionUser(ctx context.Context, tokens model.TokensInput) (*model.SessionUser, error) {
	// TODO: Sec: Validation, Auth, Permission.
	id, err := auth.GetSessionUserID(auth.Tokens(tokens))

	if err != nil {
		return &model.SessionUser{}, fmt.Errorf("unable to authenticate user: %v", err)
	}

	return &model.SessionUser{
		ID:     id.String(),
		Tokens: &model.Tokens{},
	}, nil
}
