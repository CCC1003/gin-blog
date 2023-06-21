package flag

import (
	"Blog/global"
	"Blog/models"
	"fmt"
)

func Makemigrations() {
	fmt.Println("开始迁移表结构")
	var err error
	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	// 生成表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.AdvertModel{},
		&models.UserModel{},
		&models.CommentModel{},
		&models.ArticleModel{},
		&models.BannerModel{},
		&models.MessageModel{},
		&models.ArticleModel{},
		&models.FadeBackModel{},
		&models.TagModel{},
		&models.MenuModel{},
		&models.MenuBannerModel{},
		&models.LoginDataModel{},
	)
	if err != nil {
		global.Logger.Error("生成数据表结构失败")
		return
	}
	global.Logger.Infof("生成表结构成功！")

}
