package crud

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gigamono/gigamono-api/internal/graphql/model"
	"github.com/gigamono/gigamono/pkg/configs"
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gigamono/gigamono/pkg/database/models/resource"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/files"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/services/graphql/response"
	"github.com/gigamono/gigamono/pkg/services/rest/middleware"
	"github.com/gigamono/gigamono/pkg/strs"
	"github.com/gofrs/uuid"
)

// CreateIntegration creates a new integration in the database.
func CreateIntegration(ctx context.Context, app *inits.App, integrationInput *model.IntegrationInput) (*model.Integration, error) {
	// TODO: Sec: Validation. Permission.
	userID := ctx.Value(middleware.SessionDataKey).(middleware.SessionData).User.ID

	// TODO: Validate integration config.
	integrationConfig, err := configs.NewIntegrationConfig(integrationInput.Specification, configs.JSON)
	if err != nil {
		return nil, response.Error(ctx, err.Error())
	}

	// TODO: figure workspace id // Generate obfuscated filepath.
	filePath, err := files.GenerateObfuscatedFilePath("json", uuid.Nil, "integration", nil)
	if err != nil {
		panic(errs.NewSystemError("", "generating integration spec obfuscated filename", err))
	}

	// Save integration to a file.
	if _, err = app.Filestore.Project.WriteToFile(filePath, []byte(integrationInput.Specification)); err != nil {
		panic(errs.NewSystemError("", "writing integration spec to file", err))
	}

	// Create the integration in db.
	integration := resource.Integration{Name: integrationConfig.Metadata.Name, CreatorID: userID, SpecificationFileURL: filePath}
	if err = integration.Create(&app.DB); err != nil {
		panic(errs.NewSystemError("", "creating integration", err))
	}

	return CopyIntegration(&integration), nil
}

// UploadIntegrationAvatar uploads the avatar of an associated integration.
func UploadIntegrationAvatar(_ context.Context, app *inits.App, integrationID *string, file *graphql.Upload) (*string, error) {
	// TODO: Sec: Validation. Permission.
	// Get file extension.
	ext := strings.TrimPrefix(filepath.Ext(file.Filename), ".")

	// TODO: figure workspace id // Generate obfuscated filepath.
	avatarPath, err := files.GenerateObfuscatedFilePath(ext, uuid.Nil, "avatar", nil)
	if err != nil {
		panic(errs.NewSystemError("", "generating integration spec obfuscated filename", err))
	}

	// Get uploaded file bytes.
	avatarBytes, err := ioutil.ReadAll(file.File)
	if err != nil {
		panic(errs.NewSystemError("", "reading avatar bytes", err))
	}

	// Save avatar image file.
	if _, err := app.Filestore.Image.WriteToFile(avatarPath, avatarBytes); err != nil {
		panic(errs.NewSystemError("", "writing integration avatar to file", err))
	}

	integrationUUID, err := uuid.FromString(*integrationID)
	if err != nil {
		panic(err)
	}

	// Get the integration from db.
	integration := resource.Integration{Base: models.Base{ID: integrationUUID}}
	if err = integration.GetByID(&app.DB); err != nil {
		panic(errs.NewSystemError("", "getting integration", err))
	}

	// Read specification file.
	specificationString, err := app.Filestore.Project.ReadFile(integration.SpecificationFileURL)
	if err != nil {
		panic(errs.NewSystemError("", "reading integration spec", err))
	}

	// Parse spec file.
	specificationConfig, err := configs.NewIntegrationConfig(specificationString, configs.JSON)
	if err != nil {
		panic(errs.NewSystemError("", "parsing integration spec", err))
	}

	// Delete old image file if one exists.
	if specificationConfig.Metadata.AvatarURL != nil {
		oldAvatarPath := strings.TrimPrefix(
			*specificationConfig.Metadata.AvatarURL,
			app.Config.Filestore.Image.Paths.Public+"/",
		)
		if _, err := app.Filestore.Image.DeleteFile(oldAvatarPath); err != nil {
			panic(errs.NewSystemError("", "deleting old integration avatar", err))
		}
	}

	// Update the spec avatar url.
	specificationConfig.Metadata.AvatarURL = strs.GetAddress(app.Filestore.Image.GetPublicPath(avatarPath))

	// Generate json string from updated config.
	specificationJSON, err := specificationConfig.JSON()
	if err != nil {
		panic(errs.NewSystemError("", "converting integration spec back to string", err))
	}

	// Overwrite spec file.
	if _, err := app.Filestore.Project.WriteToFile(integration.SpecificationFileURL, []byte(specificationJSON)); err != nil {
		panic(errs.NewSystemError("", "writing integration spec to file", err))
	}

	return new(string), nil
}

// GetIntegration gets an existing integration from the database.
func GetIntegration(_ context.Context, app *inits.App, integrationID *string) (*model.Integration, error) {
	// TODO: Sec: Validation, Permission.
	integrationUUID, err := uuid.FromString(*integrationID)
	if err != nil {
		panic(err)
	}

	// Get the integration from db.
	integration := resource.Integration{Base: models.Base{ID: integrationUUID}}
	if err = integration.GetByID(&app.DB); err != nil {
		panic(errs.NewSystemError("", "getting integration", err))
	}

	return CopyIntegration(&integration), nil
}

// CopyIntegration copies integration from one struct to another.
func CopyIntegration(integration *resource.Integration) *model.Integration {
	return &model.Integration{
		ID:                   integration.ID.String(),
		Name:                 integration.Name,
		CreatorID:            integration.CreatorID.String(),
		SpecificationFileURL: integration.SpecificationFileURL,
	}
}
