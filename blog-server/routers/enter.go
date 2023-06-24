package routers

import (
	"Blog/global"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	//系统配置api
	routerGroupApp.SettingsRouter()
	//图片管理api
	routerGroupApp.ImagesRouter()
	//广告管理api
	routerGroupApp.AdvertRouter()
	//菜单管理
	routerGroupApp.MenuRouter()
	//用户管理
	routerGroupApp.UserRouter()
	return router
}
