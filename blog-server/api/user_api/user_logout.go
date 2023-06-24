package user_api

import (
	"Blog/global"
	"Blog/models/res"
	"Blog/service"
	"Blog/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	token := c.Request.Header.Get("token")
	err := service.ServiceApp.UserService.Logout(claims, token)
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("用户注销成功", c)
		return
	}
	res.OkWithMessage("用户注销成功", c)

}
