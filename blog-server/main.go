package main

import (
	"Blog/core"
	_ "Blog/docs"
	"Blog/flag"
	"Blog/global"
	"Blog/routers"
)

// @title gvb_server API文档
// @version 1.0
// @description API文档
// @host 127.0.0.01:8080
// @BasePath /

func main() {
	//读取配置
	core.InitConf()
	//初始化日志
	global.Logger = core.InitLogger("logrus/log", "cong")
	//连接数据库
	global.DB = core.InitGorm()
	//连接redis
	global.Redis = core.ConnectRedis()
	//命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	//初始化路由
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Logger.Infof("gvb-server运行在：%s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Logger.Fatalf(err.Error())
	}
}
