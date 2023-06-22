package menu_api

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/res"
	"github.com/gin-gonic/gin"
)

// MenuDetailView 菜单详情
// @Tags 菜单管理
// @Summary 菜单详情
// @Description 菜单详情
// @Param data body int	true	"菜单详情"
// @Router /api/menus/:id [get]
// @Produce json
// @Success 200 {object} res.Response{data=menu_api.MenuResponse}
func (MenuApi) MenuDetailView(c *gin.Context) {
	// 先查菜单
	id := c.Param("id")

	var menuModel models.MenuModel
	err := global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	// 查连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", id)
	var banners = make([]Banner, 0)
	for _, banner := range menuBanners {
		if menuModel.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}
	MenuResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	res.OkWithData(MenuResponse, c)
	return
}
