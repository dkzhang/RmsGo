package userTempDM

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dkzhang/RmsGo/datebaseCommon/redisOps"
	"github.com/dkzhang/RmsGo/myUtils/encryptPassword"
	"github.com/dkzhang/RmsGo/myUtils/genPasswd"
	"github.com/dkzhang/RmsGo/webapi/dataManagement/userTempDM/config"
	"github.com/dkzhang/RmsGo/webapi/dataManagement/userTempDM/security"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RedisAndJwt struct {
	TheRedis         *redisOps.Redis
	TheLoginConfig   config.LoginConfig
	TheLoginSecurity security.LoginSecurity
}

func NewRedisAndJwt(redis *redisOps.Redis, cfg config.LoginConfig, se security.LoginSecurity) RedisAndJwt {
	return RedisAndJwt{
		TheRedis:         redis,
		TheLoginConfig:   cfg,
		TheLoginSecurity: se,
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
	genPasswd.RandomSeed()
	passwd = genPasswd.GeneratePasswd(rj.TheLoginConfig.ThePasswordConfig.PwLen,
		rj.TheLoginConfig.ThePasswordConfig.PwType)

	hPasswd, err := encryptPassword.GenerateFromPassword(passwd)
	if err != nil {
		return "", fmt.Errorf("encryptPassword.GenerateFromPassword error: %v", err)
	}

	err = rj.TheRedis.Set(userPasswordKey(userID), hPasswd, rj.TheLoginConfig.ThePasswordConfig.Expire)
	if err != nil {
		return "", fmt.Errorf("redis set hPasswd error: %v", err)
	}

	return passwd, nil
}

func (rj RedisAndJwt) ValidatePassword(userID int, passwd string) bool {
	if rj.TheRedis.IsExist(userPasswordKey(userID)) == false {
		return false
	}

	hPasswd := rj.TheRedis.Get(userPasswordKey(userID)).(string)

	return encryptPassword.CompareHashAndPassword(hPasswd, passwd)
}

func (rj RedisAndJwt) DelPassword(userID int) {
	rj.TheRedis.Delete(userPasswordKey(userID))
	return
}

func userPasswordKey(userID int) string {
	return fmt.Sprintf("user_passwd_%d", userID)
}

/////////////////////////////////////////////////////////////////
func (rj RedisAndJwt) CreateToken(userID int) (token string, err error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(rj.TheLoginConfig.TheTokenConfig.Expire).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err = at.SignedString([]byte(rj.TheLoginSecurity.TheTokenSecurity.Key))
	if err != nil {
		return "", err
	}

	err = rj.TheRedis.Set(userTokenKey(userID), userID, rj.TheLoginConfig.TheTokenConfig.Expire)
	if err != nil {
		return "", fmt.Errorf("redis set token error: %v", err)
	}

	return token, nil
}

func (rj RedisAndJwt) ValidateToken(r *http.Request) (userID int, err error) {
	userID, err = rj.extractTokenMetadata(extractToken(r))
	if err != nil {
		return -1, fmt.Errorf("extractTokenMetadata error: %v", err)
	}

	//check from redis
	if rj.TheRedis.IsExist(userTokenKey(userID)) == false {
		return -1, fmt.Errorf("token in redis expired: %v", err)
	}

	return userID, nil
}

func (rj RedisAndJwt) DeleteToken(userID int) {
	rj.TheRedis.Delete(userTokenKey(userID))
}

func userTokenKey(userID int) string {
	return fmt.Sprintf("user_token_%d", userID)
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (rj RedisAndJwt) extractTokenMetadata(tokenString string) (userID int, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(rj.TheLoginSecurity.TheTokenSecurity.Key), nil
	})
	if err != nil {
		return -1, fmt.Errorf("VerifyToken jwt.Parse token error: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, err = strconv.Atoi(fmt.Sprintf("%.f", claims["user_id"]))
		if err != nil {
			return -1, fmt.Errorf("extract userID error: %v", err)
		}
		return userID, nil
	}
	return -1, err
}
