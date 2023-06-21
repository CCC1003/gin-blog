package settings_api

import (
	"Blog/config"
	"Blog/core"
	"Blog/global"
	"Blog/models/res"
	"github.com/gin-gonic/gin"
)

// SettingsInfoUpdateView 修改某一项的配置信息
// @Tags 系统配置
// @Summary 修改配置信息
// @Description 修改配置信息
// @Param data query SettingsUri	false	"某一项信息的对应参数名"
// @Router /api/settings/:name [put]
// @Produce json
// @Success 200 {object} res.Response{data=string} "数据类型应为对应数据的结构体"
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var uri SettingsUri
	err := c.ShouldBindUri(&uri)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}
	switch uri.Name {
	case "site":
		var info config.SiteInfo
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.SiteInfo = info
	case "email":
		var info config.Email
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.Email = info
	case "qq":
		var info config.QQ
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.QQ = info
	case "qiniu":
		var info config.QiNiu
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.QiNiu = info
	case "jwt":
		var info config.Jwy
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.Jwy = info
	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}
	core.SetYaml()
	res.OkWithMessage("修改成功", c)
}
