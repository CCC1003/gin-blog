package settings_api

import (
	"Blog/global"
	"Blog/models/res"
	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView 显示某一项的配置信息
// @Tags 系统配置
// @Summary 配置信息
// @Description 配置信息
// @Param data query SettingsUri	false	"某一项信息的对应参数名"
// @Router /api/settings/:name [get]
// @Produce json
// @Success 200 {object} res.Response{data=string} "数据类型应为对应数据的结构体"
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var uri SettingsUri
	err := c.ShouldBindUri(&uri)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}
	global.Logger.Info(uri.Name)
	switch uri.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OkWithData(global.Config.Jwy, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}
}
