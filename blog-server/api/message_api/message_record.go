package message_api

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/res"
	"Blog/utils/jwts"
	"github.com/gin-gonic/gin"
)

type MessageRecordRequest struct {
	UserID uint `json:"user_id" binding:"required" msg:"请输入查询的用户id"`
}

func (MessageApi) MessageRecordView(c *gin.Context) {
	var cr MessageRecordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var _messageList []models.MessageModel
	var messageList = make([]models.MessageModel, 0)
	global.DB.Order("created_at asc").
		Find(&_messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)
	for _, model := range _messageList {
		//判断是一组的条件
		//send_user_id 和 rev_user_id 其中一个
		//1 2 2 1
		//1 3  3 1
		if model.RevUserID == cr.UserID || model.SendUserID == cr.UserID {
			messageList = append(messageList, model)
		}
	}
	res.OkWithData(messageList, c)
}
