package models

import (
	"Blog/global"
	"Blog/models/ctype"
	"gorm.io/gorm"
	"os"
)

// Banner Model banner表
type BannerModel struct {
	MODEL
	Path      string          `json:"path"`                        //图片路径
	Hash      string          `json:"hash"`                        //图片的hash值，用于判断重复图片
	Name      string          `gorm:"size:30" json:"name"`         //图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` //图片类型，本地还是七牛
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		//删除本地图片
		err = os.Remove(b.Path)
		if err != nil {
			global.Logger.Error(err)
			return err
		}
	}
	return nil
}
