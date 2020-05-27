package userDB

import (
	"github.com/dkzhang/RmsGo/webapi/model/user"
)

type UserDB interface {
	QueryUserByName(username string) (userInfo user.UserInfo, err error)
	QueryUserByID(userID int) (userInfo user.UserInfo, err error)
	GetAllUserInfo() (users []user.UserInfo, err error)
	UpdateUser(ui user.UserInfo) (err error)
	UpdateUserDepartment(oldDepCode string, newDep string, newDepCode string) (err error)
	InsertUser(ui user.UserInfo) (err error)
	DeleteUser(userID int) (err error)
}
