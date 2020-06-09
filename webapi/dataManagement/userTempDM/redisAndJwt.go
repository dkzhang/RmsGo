package userTempDM

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dkzhang/RmsGo/datebaseCommon/redisOps"
	"github.com/dkzhang/RmsGo/myUtils/encryptPassword"
	"github.com/dkzhang/RmsGo/myUtils/genPasswd"
	"os"
	"time"
)

type RedisAndJwt struct {
	TheRedis       *redisOps.Redis
	TheLoginConfig LoginConfig
}

func NewRedisAndJwt(redis *redisOps.Redis, cfg LoginConfig) RedisAndJwt {
	return RedisAndJwt{
		TheRedis:       redis,
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
func (rj RedisAndJwt) SetPassword(userID int) (passwd string, err error) {
	passwd = genPasswd.GeneratePasswd(rj.TheLoginConfig.ThePasswordConfig.PwLen,
		rj.TheLoginConfig.ThePasswordConfig.PwType)

	hPasswd, err := encryptPassword.GenerateFromPassword(passwd)
	if err != nil {
		return "", fmt.Errorf("encryptPassword.GenerateFromPassword error: %v", err)
	}

	err = rj.TheRedis.Set(UserPasswordKey(userID), hPasswd, rj.TheLoginConfig.ThePasswordConfig.Expire)
	if err != nil {
		return "", fmt.Errorf("redis set hPasswd error: %v", err)
	}

	return passwd, nil
}

func (rj RedisAndJwt) ValidatePassword(userID int, passwd string) bool {
	if rj.TheRedis.IsExist(UserPasswordKey(userID)) == false {
		return false
	}

	hPasswd := rj.TheRedis.Get(UserPasswordKey(userID)).(string)

	return encryptPassword.CompareHashAndPassword(hPasswd, passwd)
}

func UserPasswordKey(userID int) string {
	return fmt.Sprintf("user_passwd_%d", userID)
}

/////////////////////////////////////////////////////////////////
func (rj RedisAndJwt) CreateToken(userID int) (token string, err error) {
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(rj.TheLoginConfig.TheTokenConfig.Expire).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}

	err = rj.TheRedis.Set(UserTokenKey(userID), userID, rj.TheLoginConfig.TheTokenConfig.Expire)
	if err != nil {
		return "", fmt.Errorf("redis set token error: %v", err)
	}

	return token, nil
}

func (rj RedisAndJwt) ValidateToken(token string) (userID int, err error) {

}

func (rj RedisAndJwt) DeleteToken(userID int) {
	rj.TheRedis.Delete(UserTokenKey(userID))
}

func UserTokenKey(userID int) string {
	return fmt.Sprintf("user_token_%d", userID)
}
