package routers

import (
	"Blog/api"
)

func (router RouterGroup) SettingsRouter() {
	app := api.ApiGroupApp.SettingsApi
	router.GET("settings/:name", app.SettingsInfoView)
	router.PUT("settings/:name", app.SettingsInfoUpdateView)
}
