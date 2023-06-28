package tag_api

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//重复判断
	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		res.FailWithMessage("标签不存在", c)
	}
	maps := structs.Map(&cr)
	err = global.DB.Debug().Model(&tag).Updates(maps).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("修改标签失败", c)
	}
	res.OkWithMessage("修改标签成功", c)
}
