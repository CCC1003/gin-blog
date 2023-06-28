package routers

import (
	"Blog/api"
	"Blog/middleware"
)

func (router RouterGroup) MessageRouter() {
	app := api.ApiGroupApp.MessageApi
	router.POST("messages", app.MessageCreateView)
	router.GET("message_all", app.MessageListAllView)
	router.GET("messages", middleware.JwtAuth(), app.MessageListView)
	router.GET("messages_record", middleware.JwtAuth(), app.MessageRecordView)

}
