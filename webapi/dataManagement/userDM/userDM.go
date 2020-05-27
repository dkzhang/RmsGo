package userDM

import "github.com/dkzhang/RmsGo/webapi/model/user"

func GetInstance() (udm UserDM, err error) {

	return nil, nil
}

type UserDM interface {
	QueryUserByName(userName string) (user.UserInfo, error)
	QueryUserByID(userID int) (user.UserInfo, error)
	QueryUserByDepartmentCode(dc string) ([]user.UserInfo, error)
}
