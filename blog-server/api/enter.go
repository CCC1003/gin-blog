package api

import (
	"Blog/api/advert_api"
	"Blog/api/images_api"
	"Blog/api/menu_api"
	"Blog/api/settings_api"
	"Blog/api/tag_api"
	"Blog/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImageApi    images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
	UserApi     user_api.UserApi
	TagApi      tag_api.TagApi
}

var ApiGroupApp = new(ApiGroup)
