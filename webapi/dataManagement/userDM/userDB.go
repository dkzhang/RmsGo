package userDM

import (
	"github.com/dkzhang/RmsGo/webapi/model/user"
)

type UserDB interface {
	QueryUserByName(username string) (userInfo user.UserInfo, err error)
	QueryUserByID(userID int) (userInfo user.UserInfo, err error)
}
