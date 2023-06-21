package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = 1
	SignGitee SignStatus = 2
	SignEmail SignStatus = 3
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "QQ"
	case SignEmail:
		str = "邮箱"
	case SignGitee:
		str = "gitee"
	default:
		str = "其他"
	}
	return str
}
