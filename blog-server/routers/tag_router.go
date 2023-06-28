package routers

import "Blog/api"

func (router RouterGroup) TagRouter() {
	app := api.ApiGroupApp.TagApi
	router.POST("tags", app.TagCreateView)
	router.GET("tags", app.TagListView)
	router.PUT("tags/:id", app.TagUpdateView)
	router.DELETE("tags", app.TagRemoveView)
}
