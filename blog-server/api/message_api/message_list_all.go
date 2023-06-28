package message_api

import (
	"Blog/models"
	"Blog/models/res"
	"Blog/service/common"
	"github.com/gin-gonic/gin"
)

func (MessageApi) MessageListAllView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, _ := common.ComList(models.MessageModel{}, common.Option{
		PageInfo: cr,
	})
	res.OkWithList(list, count, c)
}
