package user_api

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/ctype"
	"Blog/models/res"
	"github.com/gin-gonic/gin"
)

type UserRole struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	NickName string     `json:"nick_name"` //防止用户昵称非法，管理员有能力修改
	UserID   uint       `json:"user_id" binding:"required" msg:"用户id错误"`
}

// UserUpdateRoleView 用户权限变更
// @Tags 用户管理
// @Summary 用户权限变更
// @Description 用户权限变更
// @Param data query UserRole 	false	"用户角色、昵称、id"
// @Router /api/user_role [put]
// @Produce json
// @Success 200 {object} res.Response{data=string} ""
func (UserApi) UserUpdateRoleView(c *gin.Context) {
	var cr UserRole
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err = global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		res.FailWithMessage("用户id错误，用户不存在", c)
		return
	}
	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("修改权限失败", c)
		return
	}
	res.FailWithMessage("修改权限成功", c)
}
