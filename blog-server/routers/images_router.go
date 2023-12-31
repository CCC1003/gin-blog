package routers

import (
	"Blog/api"
)

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImageApi
	router.GET("images", app.ImageListView)
	router.GET("image_names", app.ImageNameListView)
	router.POST("images", app.ImageUploadView)
	router.DELETE("images", app.ImageRemoveView)
	router.PUT("images", app.ImageUpdateView)
}
