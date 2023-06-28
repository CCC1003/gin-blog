package user_ser

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/ctype"
	"Blog/utils/pwd"
	"errors"
)

const Avatar = "/uploads/file/default.jpg"

func (UserService) CreateUser(userName, nickName, password string, role ctype.Role, email string, ip string) error {
	//判断用户名是否存在
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error

	if err == nil {
		return errors.New("用户名已存在")
	}
	//对密码进行hash
	hashAndSalt := pwd.HashAndSalt(password)

	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashAndSalt,
		Email:      email,
		Role:       role,
		AvatarID:   Avatar,
		IP:         ip,
		Addr:       "内网地址",
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
