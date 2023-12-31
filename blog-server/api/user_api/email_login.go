package user_api

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/res"
	"Blog/utils/jwts"
	"Blog/utils/pwd"
	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

// EmailLoginView  用户邮箱登录
// @Tags 用户管理
// @Summary 用户登录
// @Description 用户登录
// @Param data query EmailLoginRequest	false	"用户名和密码"
// @Router /api/email_login [post]
// @Produce json
// @Success 200 {object} res.Response{data=jwts.JwtPayLoad} ""
func (UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ?", cr.UserName).Error
	if err != nil {
		global.Logger.Warn("用户名不存在")
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	//校验密码
	isCheck := pwd.ComparePasswords(userModel.Password, cr.Password)
	if !isCheck {
		global.Logger.Warn("用户名密码错误")
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	//登录成功，生成Token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("token生成失败", c)
		return
	}
	res.OkWithData(token, c)

}
