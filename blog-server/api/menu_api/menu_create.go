package menu_api

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/ctype"
	"Blog/models/res"
	"github.com/gin-gonic/gin"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type menuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`                //切换的时间
	BannerTime    int         `json:"banner_time" structs:"banner_time"`                    //切换的时间
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号" structs:"sort"` //菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`                          //具体图片的顺序
}

// MenuCreateView 添加菜单
// @Tags 菜单管理
// @Summary 添加菜单
// @Description 添加菜单
// @Param data body menuRequest	true	"添加菜单"
// @Router /api/menus [post]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr menuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//菜单的重复值判断
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, "title = ? or path = ?", cr.Title, cr.Path).RowsAffected
	if count > 0 {
		res.FailWithMessage("重复的菜单", c)
		return
	}
	//创建banner数据入库
	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("菜单添加失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		res.FailWithMessage("菜单添加成功", c)
		return
	}
	var menuBannerList []models.MenuBannerModel
	for _, imageSort := range cr.ImageSortList {
		//判断image_id是否真正有这张图片
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: imageSort.ImageID,
			Sort:     imageSort.Sort,
		})
	}
	//给第三张表入库
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("菜单图片关联失败", c)
		return
	}
	res.OkWithMessage("菜单添加成功", c)
}
