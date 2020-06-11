package userDM

import "github.com/dkzhang/RmsGo/webapi/model/user"

type UserDM interface {
	QueryUserByName(userName string) (user.UserInfo, error)
	QueryUserByID(userID int) (user.UserInfo, error)

	QueryUserByDepartmentCode(dc string) []user.UserInfo
	QueryUserByFilter(userFilter func(user.UserInfo) bool) []user.UserInfo

	IsUserNameExist(userName string) bool

	UpdateUserPreCheck(userNew user.UserInfo) (string, error)
	UpdateUser(userNew user.UserInfo) (err error)

	InsertUserPreCheck(userNew user.UserInfo) (string, error)
	InsertUser(userNew user.UserInfo) (err error)

	DeleteUser(userID int) (err error)
}
