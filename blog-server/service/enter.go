package service

import (
	"Blog/service/image_ser"
	"Blog/service/user_ser"
)

type ServiceGroup struct {
	ImageService image_ser.ImageService
	UserService  user_ser.UserService
}

var ServiceApp = new(ServiceGroup)
