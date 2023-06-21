package api

import (
	"Blog/api/advert_api"
	"Blog/api/images_api"
	"Blog/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImageApi    images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
}

var ApiGroupApp = new(ApiGroup)
