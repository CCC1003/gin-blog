package menu_api

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// MenuRemoveView 菜单删除
// @Tags 菜单管理
// @Summary 菜单删除
// @Description 菜单删除
// @Param data body models.RemoveRequest	true	"删除菜单"
// @Router /api/menus [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	//事务
	global.DB.Transaction(func(tx *gorm.DB) error {
		if err != nil {
			global.Logger.Error(err)
			return err
		}
		err = global.DB.Delete(&menuList).Error
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
