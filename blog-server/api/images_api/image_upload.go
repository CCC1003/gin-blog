package images_api

import (
	"Blog/global"
	"Blog/models/res"
	"Blog/service"
	"Blog/service/image_ser"
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
)

// ImageUploadView 上传单个图片，返回图片的url
// @Tags 图片管理
// @Summary 图片上传
// @Description 图片上传
// @Param data body string	true	"本应使用表单多图片文件参数"
// @Router /api/images [post]
// @Produce json
// @Success 200 {object} res.Response{data=image_ser.FileUploadResponse[models.BannerModel]}
func (ImagesApi) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
	}
	//判断路径是否存在
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		//不存在就创建
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Logger.Error(err)
		}
	}
	var resList []image_ser.FileUploadResponse

	for _, file := range fileList {
		//上传文件
		serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		if !global.Config.QiNiu.Enable {
			//本地保存
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Logger.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}
		resList = append(resList, serviceRes)
	}
	res.OkWithData(resList, c)
}
