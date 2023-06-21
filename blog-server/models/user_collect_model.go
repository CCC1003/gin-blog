package models

import "time"

//UserCollectModel 自定义第三张表 记录用户什么时候收藏了文章

type UserCollectModel struct {
	UserID    uint      `gorm:"primaryKey"`        //用户ID
	UserModel UserModel `gorm:"foreignKey:UserID"` //用户

	ArticleID    uint         `gorm:"primaryKey"`           //文章ID
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"` //文章
	CreateAt     time.Time    //创建时间
}
