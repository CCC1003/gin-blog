package image_ser

import (
	"Blog/global"
	"Blog/models"
	"Blog/models/ctype"
	"Blog/plugins/qiniu"
	"Blog/utils"
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

var (
	//图片上传的白名单
	WhiteImageList = []string{
		"jpg",
		"png",
		"jpeg",
		"ico",
		"tiff",
		"gif",
		"svg",
		"webp",
	}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"` //图片是否上传成功
	Msg       string `json:"msg"`
}

// 文件上传的方法
func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fileName := file.Filename
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, fileName)
	res.FileName = filePath

	//文件白名单判断
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = "非法文件"
		return
	}
	//判断文件名大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超过设定大小，设定大小为：%dMB", global.Config.Upload.Size)
	}

	//读取文件内容 hash
	fileObj, err := file.Open()
	if err != nil {
		global.Logger.Error(err)
	}
	byteData, err := io.ReadAll(fileObj)
	if err != nil {
		global.Logger.Error(err)
	}
	imageHash := utils.Md5(byteData)
	//在数据库中判断文件是否存在
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	if err == nil {
		res.Msg = "图片已存在"
		res.FileName = bannerModel.Path
		return
	}
	fileType := ctype.Local
	res.Msg = "图片上传本地成功"
	res.IsSuccess = true
	if global.Config.QiNiu.Enable {
		filePath, err = qiniu.UploadImage(byteData, fileName, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Logger.Error(err)
			res.Msg = err.Error()
			return
		}
		res.FileName = filePath
		res.Msg = "图片上传七牛云成功"
		fileType = ctype.QiNiu
	}
	//图片入库
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Name:      fileName,
		Hash:      imageHash,
		ImageType: fileType,
	})
	return
}
