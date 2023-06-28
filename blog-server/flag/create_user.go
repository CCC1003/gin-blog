package flag

import (
	"Blog/global"
	"Blog/models/ctype"
	"Blog/service/user_ser"
	"fmt"
)

func CreateUser(permissions string) {
	//用户名 昵称 密码 确认密码 邮箱
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)
	fmt.Printf("请输入用户名：")
	fmt.Scan(&userName)
	fmt.Printf("请输入昵称：")
	fmt.Scan(&nickName)
	fmt.Printf("请输入密码：")
	fmt.Scan(&password)
	fmt.Printf("请再次输入密码：")
	fmt.Scan(&rePassword)
	fmt.Printf("请输入邮箱：")
	fmt.Scan(&email)
	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}
	err := user_ser.UserService{}.CreateUser(userName, nickName, password, role, email, "")
	if err != nil {
		global.Logger.Error(err)
		return
	}

	global.Logger.Infof("用户%s创建成功", userName)
}
