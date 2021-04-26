package mockdata

import (
	"github.com/gigamono/gigamono-api/internal/mockql/model"
	"github.com/gigamono/gigamono/pkg/strs"
)

func getUserProfile() *model.Profile {
	return &model.Profile{
		Username:    "nypro",
		Email:       "smartnypro@gmail.com",
		FirstName:   strs.GetAddress("Smart"),
		LastName:    strs.GetAddress("Nypro"),
		Avatar32url: strs.GetAddress("avatars/usr-007ddc8b-ccd2-42d1-9c29-b6c2f9114829/avatar-32-nypro.png"),
	}
}
