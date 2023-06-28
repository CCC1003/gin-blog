package message_api

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/res"
	"github.com/gin-gonic/gin"
)

type MessageRequest struct {
	SendUserID uint   `gorm:"primaryKey" json:"send_user_id" binding:"required"` //发送人id
	RevUserID  uint   `gorm:"primaryKey" json:"rev_user_id" binding:"required"`  //接收人id
	Content    string `json:"content" binding:"required"`                        //消息内容
}

// MessageCreateView  发布消息
func (MessageApi) MessageCreateView(c *gin.Context) {
	//当前用户发布消息
	//SendUserID 就是当前登录人的id
	var cr MessageRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var senUser, recvUser models.UserModel
	err = global.DB.Take(&senUser, cr.SendUserID).Error
	if err != nil {
		res.FailWithMessage("发送人不存在", c)
		return
	}
	err = global.DB.Take(&recvUser, cr.RevUserID).Error
	if err != nil {
		res.FailWithMessage("接收人不存在", c)
		return
	}
	err = global.DB.Create(&models.MessageModel{
		SendUserID:       cr.SendUserID,
		SendUserNickName: senUser.NickName,
		SendUserAvatar:   senUser.AvatarID,
		RevUserID:        cr.RevUserID,
		RevUserNickName:  recvUser.NickName,
		RevUserAvatar:    recvUser.NickName,
		IsRead:           false,
		Content:          cr.Content,
	}).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("消息发送失败", c)
		return
	}
	res.OkWithMessage("消息发送成功", c)
	return
}
