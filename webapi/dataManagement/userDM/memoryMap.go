package userDM

import (
	"fmt"
	userDB2 "github.com/dkzhang/RmsGo/webapi/dataManagement/userDB"
	"github.com/dkzhang/RmsGo/webapi/model/user"
)

type MemoryMap struct {
	userInfoByID   map[int]*user.UserInfo
	userInfoByName map[string]*user.UserInfo
	userDB         userDB2.UserDB
}

func NewMemoryMap(udb userDB2.UserDB) (nmm MemoryMap, err error) {
	nmm.userDB = udb

	users, err := nmm.userDB.GetAllUserInfo()
	if err != nil {
		return MemoryMap{},
			fmt.Errorf("generate new MemoryMap failed since GetAllUserInf error: %v", err)
	}

	nmm.userInfoByID = make(map[int]*user.UserInfo, len(users))
	nmm.userInfoByName = make(map[string]*user.UserInfo, len(users))

	for _, v := range users {
		user := v //Create a temp variable <user> here is very necessary
		nmm.userInfoByID[v.UserID] = &user
		nmm.userInfoByName[v.UserName] = &user
	}
	return nmm, nil
}

func (udm *MemoryMap) QueryUserByName(userName string) (user.UserInfo, error) {
	if v, ok := udm.userInfoByName[userName]; ok {
		return *v, nil
	} else {
		return user.UserInfo{}, fmt.Errorf("user (name = %s) not exist", userName)
	}
}

func (udm *MemoryMap) QueryUserByID(userID int) (user.UserInfo, error) {
	if v, ok := udm.userInfoByID[userID]; ok {
		return *v, nil
	} else {
		return user.UserInfo{}, fmt.Errorf("user (id = %d) not exist", userID)
	}
}

func (udm *MemoryMap) QueryUserByDepartmentCode(dc string) ([]user.UserInfo, error) {
	return udm.QueryUserByFilter(func(userInfo user.UserInfo) bool {
		return userInfo.DepartmentCode == dc
	})
}

func (udm *MemoryMap) QueryUserByFilter(userFilter func(user.UserInfo) bool) ([]user.UserInfo, error) {
	uis := make([]user.UserInfo, 0)
	for _, v := range udm.userInfoByID {
		if userFilter(*v) == true {
			uis = append(uis, *v)
		}
	}
	if len(uis) == 0 {
		return nil, fmt.Errorf("no user found with giving filter")
	} else {
		return uis, nil
	}
}

func (udm *MemoryMap) IsUserNameExist(userName string) bool {
	_, ok := udm.userInfoByName[userName]
	return ok
}
func (udm *MemoryMap) UpdateUser(userNew user.UserInfo) (err error) {
	//check exist
	userOld, ok := udm.userInfoByID[userNew.UserID]
	if ok == false {
		return fmt.Errorf("UpdateUser failed since user (userId = %d) not found", userNew.UserID)
	}

	//check duplicate name
	if userNew.UserName != userOld.UserName {
		if udm.IsUserNameExist(userNew.UserName) {
			return fmt.Errorf("UpdateUser failed since userName <%s> is already exist", userNew.UserName)
		}
	}

	//update user info
	//general, only 6 property could be update
	err = udm.userDB.UpdateUser(userNew)
	if err != nil {
		return fmt.Errorf("udm.userDB.UpdateUser error: %v", err)
	} else {
		//update user info in memory map
		userOld.UserName = userNew.UserName
		userOld.ChineseName = userNew.ChineseName
		userOld.Section = userNew.Section
		userOld.Mobile = userNew.Mobile
		userOld.Remarks = userNew.Remarks
		userOld.Status = userNew.Status
	}

	//special ops for approver
	if userOld.Role == user.RoleApprover {
		if userOld.Department == userNew.Department && userOld.DepartmentCode == userNew.DepartmentCode {
			//department info not modified
			return nil
		} else {
			//update all user in the same department
			err = udm.userDB.UpdateUserDepartment(userOld.DepartmentCode,
				userNew.Department, userNew.DepartmentCode)
			if err != nil {
				return fmt.Errorf("udm.userDB.UpdateUserDepartment error: %v", err)
			}

			for _, u := range udm.userInfoByID {
				if u.DepartmentCode == userOld.DepartmentCode {
					u.DepartmentCode = userNew.DepartmentCode
					u.Department = userNew.Department
				}
			}
		}
	}

	return nil
}
