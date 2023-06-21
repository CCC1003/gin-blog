package advert_api

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// AdvertUpdateView 更新广告
// @Tags 广告管理
// @Summary 更新广告
// @Description 更新广告
// @Param data body AdvertRequest	true	"广告的一些参数"
// @Router /api/adverts/:id [put]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//重复判断
	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		res.FailWithMessage("广告不存在", c)
	}
	maps := structs.Map(&cr)
	err = global.DB.Debug().Model(&advert).Updates(maps).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("修改广告失败", c)
	}
	res.OkWithMessage("修改广告成功", c)
}
