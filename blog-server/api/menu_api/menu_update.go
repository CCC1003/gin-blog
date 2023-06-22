package menu_api

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// MenuUpdateView 菜单更新
// @Tags 菜单管理
// @Summary 菜单更新
// @Description 菜单更新
// @Param data body menuRequest	true	"菜单更新"
// @Router /api/menus/:id [put]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr menuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//先把之前banner清空
	id := c.Param("id")
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	global.DB.Model(&menuModel).Association("Banners").Clear()
	//如果选择了banner就添加
	if len(cr.ImageSortList) > 0 {
		//操作第三方表
		var menuBanners []models.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			menuBanners = append(menuBanners, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&menuBanners).Error
		if err != nil {
			global.Logger.Error(err)
			res.FailWithMessage("创建菜单图片失败", c)
			return
		}
	}
	//普通更新
	maps := structs.Map(&cr)
	err = global.DB.Debug().Model(&menuModel).Updates(maps).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("修改菜单失败", c)
		return
	}
	res.OkWithMessage("修改菜单成功", c)
}
