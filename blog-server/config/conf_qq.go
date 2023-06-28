package config

type QQ struct {
	AppID    string `json:"app_id" yaml:"app_id"`
	Key      string `json:"key" yaml:"key"`
	Redirect string `json:"redirect" yaml:"redirect"` //登录之后的回调地址
}

//https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=101974593&redirect_uri=http://blog.fengfengzhidao.com/Login?flag=qq
