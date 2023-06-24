package user_ser

import (
	"Blog/service/redis_ser"
	"Blog/utils/jwts"
	"time"
)

type UserService struct {
}

func (UserService) Logout(claims *jwts.CustomClaims, token string) error {
	//需要计算距离现在的过期时间
	exp := claims.ExpiresAt //截止时间点
	now := time.Now()

	diff := exp.Time.Sub(now) //截止时间点减去现在时间点的时间段 秒数

	return redis_ser.LogoutRedis(token, diff)
}
