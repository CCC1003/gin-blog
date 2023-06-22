package menu_api

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/res"
	"github.com/gin-gonic/gin"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}
type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

// MenuListView 菜单列表
// @Tags 菜单管理
// @Summary 菜单列表
// @Description 菜单列表
// @Router /api/menus [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]MenuResponse}
func (MenuApi) MenuListView(c *gin.Context) {
	//先查菜单
	var menuList []models.MenuModel
	var menuIDList []uint
	global.DB.Find(&menuList).Select("id").Scan(&menuIDList)
	//查连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in ?", menuIDList)
	var menus []MenuResponse
	for _, model := range menuList {
		//model就是一个菜单
		var banners = make([]Banner, 0)
		for _, banner := range menuBanners {
			if model.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
		}
		menus = append(menus, MenuResponse{
			MenuModel: model,
			Banners:   banners,
		})
	}
	res.OkWithData(menus, c)
	return
}
