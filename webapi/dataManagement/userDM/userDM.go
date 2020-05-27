package userDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi"
	"github.com/dkzhang/RmsGo/webapi/model/user"
)

func GetInstance() {

}

type UserDM struct {
	userInfoByID   map[int]*user.UserInfo
	userInfoByName map[string]*user.UserInfo
}

func (udm *UserDM) init() error {
	users, err := GetAllUserInfo(webapi.TheContext.TheDb)
	if err != nil {
		return fmt.Errorf("init faild since GetAllUserInf error: %v", err)
	}

	udm.userInfoByID = make(map[int]*user.UserInfo, len(users))
	udm.userInfoByName = make(map[string]*user.UserInfo, len(users))

	for _, v := range users {
		user := v //Create a temp variable <user> here is very necessary
		udm.userInfoByID[v.UserID] = &user
		udm.userInfoByName[v.UserName] = &user
	}
	return nil
}

func (udm *UserDM) QueryUserByName(userName string) (user.UserInfo, error) {
	if v, ok := udm.userInfoByName[userName]; ok {
		return *v, nil
	} else {
		return user.UserInfo{}, fmt.Errorf("user (name = %s) not exist", userName)
	}
}

func (udm *UserDM) QueryUserByID(userID int) (user.UserInfo, error) {
	if v, ok := udm.userInfoByID[userID]; ok {
		return *v, nil
	} else {
		return user.UserInfo{}, fmt.Errorf("user (id = %d) not exist", userID)
	}
}

func (udm *UserDM) QueryUserByDepartmentCode(dc string) ([]user.UserInfo, error) {
	uis := make([]user.UserInfo, 0)
	for _, v := range udm.userInfoByID {
		if v.DepartmentCode == dc {
			uis = append(uis, *v)
		}
	}
	if len(uis) == 0 {
		return nil, fmt.Errorf("no user found with department code <%s>", dc)
	} else {
		return uis, nil
	}
}
