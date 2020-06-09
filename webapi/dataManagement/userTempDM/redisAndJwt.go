package userTempDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/datebaseCommon/redisOps"
)

type RedisAndJwt struct {
	TheRedis       *redisOps.Redis
	TheLoginConfig LoginConfig
}

func NewRedisAndJwt(r *redisOps.Redis, cfg LoginConfig) RedisAndJwt {
	return RedisAndJwt{
		TheRedis:       r,
		TheLoginConfig: cfg,
	}
}

func (rj RedisAndJwt) IsSmsLock(userID int) bool {
	return rj.TheRedis.IsExist(UserSmsKey(userID))
}

func (rj RedisAndJwt) LockSms(userID int) error {
	return rj.TheRedis.Set(UserSmsKey(userID), userID, rj.TheLoginConfig.TheSmsConfig.LockTime)
}

func UserSmsKey(userID int) string {
	return fmt.Sprintf("user_smslock_%d", userID)
}

/////////////////////////////////////////////////////////////////
func (rj RedisAndJwt) SetPassword(userID int) error {

}

func (rj RedisAndJwt) ValidatePassword(userID int) (bool, error) {

}

/////////////////////////////////////////////////////////////////
func (rj RedisAndJwt) CreateToken(userID int) (token string, err error) {

}

func (rj RedisAndJwt) ValidateToken(token string) (userID int, err error) {

}

func (rj RedisAndJwt) DeleteToken(userID int) (err error) {

}
