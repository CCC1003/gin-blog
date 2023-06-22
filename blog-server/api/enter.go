package api

import (
	"Blog/api/advert_api"
	"Blog/api/images_api"
	"Blog/api/menu_api"
	"Blog/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImageApi    images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
}

var ApiGroupApp = new(ApiGroup)
