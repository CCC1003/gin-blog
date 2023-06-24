package user_api

import (
	"Blog/models"
	"Blog/models/ctype"
	"Blog/models/res"
	"Blog/service/common"
	"Blog/utils/desens"
	"Blog/utils/jwts"
	"github.com/gin-gonic/gin"
)

// UserListView  用户列表查询
// @Tags 用户管理
// @Summary 用户列表
// @Description 用户列表
// @Param data query models.PageInfo	false	"分页参数"
// @Router /api/users [get]
// @Produce json
// @Success 200 {object} res.Response{data=models.UserModel} ""
func (UserApi) UserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
	})
	var users []models.UserModel
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			user.UserName = ""
		}
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)
		users = append(users, user)
	}
	res.OkWithList(users, count, c)
}
