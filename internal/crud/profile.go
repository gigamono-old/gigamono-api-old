package crud

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gigamono/gigamono-api/internal/graphql/model"
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gigamono/gigamono/pkg/database/models/resource"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/files"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/strs"
	"github.com/gofrs/uuid"
)

// UploadProfileAvatar uploads the avatar of an associated profile.
func UploadProfileAvatar(_ context.Context, app *inits.App, profileID *string, file *graphql.Upload) (*string, error) {
	// TODO: Sec: Validation. Permission.
	// Get file extension.
	ext := strings.TrimPrefix(filepath.Ext(file.Filename), ".")

	// TODO: figure workspace id // Generate obfuscated filepath.
	avatarPath, err := files.GenerateObfuscatedFilePath(ext, uuid.Nil, "avatar", nil)
	if err != nil {
		panic(errs.NewSystemError("", "generating profile spec obfuscated filename", err))
	}

	// Get uploaded file bytes.
	avatarBytes, err := ioutil.ReadAll(file.File)
	if err != nil {
		panic(errs.NewSystemError("", "reading avatar bytes", err))
	}

	// Save avatar image file.
	if _, err := app.Filestore.Image.WriteToFile(avatarPath, avatarBytes); err != nil {
		panic(errs.NewSystemError("", "writing profile avatar to file", err))
	}

	profileUUID, err := uuid.FromString(*profileID)
	if err != nil {
		panic(err)
	}

	// Get the profile from db.
	profile := resource.Profile{Base: models.Base{ID: profileUUID}}
	if err = profile.GetByID(&app.DB); err != nil {
		panic(errs.NewSystemError("", "getting profile", err))
	}

	// Delete old image file if one exists.
	if profile.AvatarURL != nil {
		oldAvatarPath := strings.TrimPrefix(
			*profile.AvatarURL,
			app.Config.Filestore.Image.Paths.Public+"/",
		)
		if _, err := app.Filestore.Image.DeleteFile(oldAvatarPath); err != nil {
			panic(errs.NewSystemError("", "deleting old profile avatar", err))
		}
	}

	// Update profile avatarURL.
	profile.AvatarURL = strs.GetAddress(app.Image.GetPublicPath(avatarPath))
	if profile.UpdateByID(&app.DB); err != nil {
		panic(errs.NewSystemError("", "updating profile", err))
	}

	return new(string), nil
}

// UpdateProfile updates an existing profile from the database.
func UpdateProfile(_ context.Context, app *inits.App, profileID *string, profileInput *model.ProfileInput) (*model.Profile, error) {
	// TODO: Sec: Validation, Permission.
	profileUUID, err := uuid.FromString(*profileID)
	if err != nil {
		panic(err)
	}

	// Fetch profile from the db.
	profile := resource.Profile{Base: models.Base{ID: profileUUID}}
	if err = profile.GetByID(&app.DB); err != nil {
		panic(errs.NewSystemError("", "getting profile", err))
	}

	// Update profile fields with non-empty non-nil values.
	if profileInput.Username != nil {
		profile.Username = profileInput.Username
	}

	if profileInput.FirstName != nil {
		profile.FirstName = profileInput.FirstName
	}

	if profileInput.LastName != nil {
		profile.LastName = profileInput.LastName
	}

	if profileInput.Email != nil {
		profile.Email = profileInput.Email
	}

	// Update profile in db.
	if err = profile.UpdateByID(&app.DB); err != nil {
		panic(errs.NewSystemError("", "updating profile", err))
	}

	return CopyProfile(&profile), nil
}

// CopyProfile copies profile from one struct to another.
func CopyProfile(profile *resource.Profile) *model.Profile {
	return &model.Profile{
		ID:        profile.ID.String(),
		Username:  profile.Username,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Email:     profile.Email,
		AvatarURL: profile.AvatarURL,
		UserID:    profile.UserID.String(),
	}
}
