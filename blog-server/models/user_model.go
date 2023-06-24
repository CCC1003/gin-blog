package models

import "Blog/models/ctype"

// UserModel用户表
type UserModel struct {
	MODEL
	NickName       string           `gorm:"size:36" json:"nick_name"`                                                             //昵称
	UserName       string           `gorm:"size:36" json:"user_name"`                                                             //用户名
	Password       string           `gorm:"size:128" json:"-"`                                                                    //密码
	AvatarID       string           `gorm:"size:256" json:"avatar_id"`                                                            //头像id
	Email          string           `gorm:"size:128" json:"email"`                                                                //邮箱
	Tel            string           `gorm:"size:18" json:"tel"`                                                                   //用户名
	Addr           string           `gorm:"size:64" json:"addr"`                                                                  //地址
	Token          string           `gorm:"size:64" json:"token"`                                                                 //其他平台的唯一id
	IP             string           `gorm:"size:20" json:"IP"`                                                                    //ip地址
	Role           ctype.Role       `gorm:"size:4;default:1" json:"role"`                                                         // //权限 1 管理员 2 普通用户 3 游客
	SignStatus     ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"`                                                  //注册来源
	ArticleModels  []ArticleModel   `gorm:"foreignKey:UserID" json:"-""`                                                          //发布的文章列表
	CollectsModels []ArticleModel   `gorm:"many2many:User_collect_model;joinForeignKey:UserID;JoinReferences:ArticleID" json:"-"` //收藏的文章列表
}
