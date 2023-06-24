package user_api

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/res"
	"Blog/utils"
	"Blog/utils/jwts"
	"github.com/gin-gonic/gin"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd" binding:"required"` //旧密码
	Pwd    string `json:"pwd" binding:"required"`     //新密码
}

// UserUpdatePassword 修改登录人的id
// @Tags 用户管理
// @Summary 用户更新密码
// @Description 用户更新密码
// @Param data query UpdatePasswordRequest 	false	"旧密码和新密码"
// @Router /api/user_password [put]
// @Produce json
// @Success 200 {object} res.Response{data=string} ""
func (UserApi) UserUpdatePassword(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	// 判断密码是否一致
	if !utils.ComparePasswords(user.Password, cr.OldPwd) {
		res.FailWithMessage("密码错误", c)
		return
	}
	hashPwd := utils.HashAndSalt(cr.Pwd)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("密码修改失败", c)
		return
	}
	res.OkWithMessage("密码修改成功", c)
	return
}
