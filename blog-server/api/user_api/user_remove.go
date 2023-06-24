package user_api

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (UserApi) UserRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var userList []models.UserModel
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	//事务
	global.DB.Transaction(func(tx *gorm.DB) error {
		//TODO : 删除用户，消息表，评论表，用户收藏的文章，用户发布的文章
		if err != nil {
			global.Logger.Error(err)
			return err
		}
		err = global.DB.Delete(&userList).Error
		if err != nil {
			global.Logger.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Logger.Error(err)
		res.OkWithMessage("删除菜单失败", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个菜单", count), c)
}
