package userTempDM

import "net/http"

type UserTempDM interface {
	IsSmsLock(userID int) bool
	LockSms(userID int) error

	SetPassword(userID int) (passwd string, err error)
	DelPassword(userID int)
	ValidatePassword(userID int, passwd string) bool

	CreateToken(userID int) (token string, err error)
	ValidateToken(r *http.Request) (userID int, err error)
	DeleteToken(userID int)
}
