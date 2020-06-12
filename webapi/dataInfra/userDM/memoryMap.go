package userDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/userDB"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/sirupsen/logrus"
)

type MemoryMap struct {
	userInfoByID   map[int]*user.UserInfo
	userInfoByName map[string]*user.UserInfo
	theUserDB      userDB.UserDB
}

func NewMemoryMap(udb userDB.UserDB) (nmm MemoryMap, err error) {
	nmm.theUserDB = udb

	users, err := nmm.theUserDB.GetAllUserInfo()
	if err != nil {
		return MemoryMap{},
			fmt.Errorf("generate new MemoryMap failed since GetAllUserInf error: %v", err)
	}

	logMap.Log(logMap.DEFAULT).WithFields(logrus.Fields{
		"AllUserInfo": users,
	}).Info("NewMemoryMap theUserDB.GetAllUserInfo success.")

	nmm.userInfoByID = make(map[int]*user.UserInfo, len(users))
	nmm.userInfoByName = make(map[string]*user.UserInfo, len(users))

	for _, v := range users {
		user := v //Create a temp variable <user> here is very necessary
		nmm.userInfoByID[v.UserID] = &user
		nmm.userInfoByName[v.UserName] = &user
	}

	logMap.Log(logMap.DEFAULT).WithFields(logrus.Fields{
		"userInfoByID":   nmm.userInfoByID,
		"userInfoByName": nmm.userInfoByName,
	}).Info("NewMemoryMap load data to map success.")

	return nmm, nil
}

func (udm MemoryMap) QueryUserByName(userName string) (user.UserInfo, error) {
	if user.CheckUserName(userName) == false {
		return user.UserInfo{},
			fmt.Errorf("QueryUserByName failed: username  <%s> illegal", userName)
	}

	if v, ok := udm.userInfoByName[userName]; ok {
		return *v, nil
	} else {
		return user.UserInfo{}, fmt.Errorf("user (name = %s) not exist", userName)
	}
}

func (udm MemoryMap) QueryUserByID(userID int) (user.UserInfo, error) {
	if v, ok := udm.userInfoByID[userID]; ok {
		return *v, nil
	} else {
		return user.UserInfo{}, fmt.Errorf("user (id = %d) not exist", userID)
	}
}

func (udm MemoryMap) QueryUserByDepartmentCode(dc string) []user.UserInfo {
	return udm.QueryUserByFilter(func(userInfo user.UserInfo) bool {
		return userInfo.DepartmentCode == dc
	})
}

func (udm MemoryMap) QueryUserByFilter(userFilter func(user.UserInfo) bool) []user.UserInfo {
	uis := make([]user.UserInfo, 0)
	for _, v := range udm.userInfoByID {
		if userFilter(*v) == true {
			uis = append(uis, *v)
		}
	}
	return uis
}

func (udm MemoryMap) IsUserNameExist(userName string) bool {
	_, ok := udm.userInfoByName[userName]
	return ok
}

// return "",nil if all pre-check pass.
// if pre-check failed, return msg with Chinese information for the user, return err for program.
func (udm MemoryMap) UpdateUserPreCheck(userNew user.UserInfo) (string, error) {
	// check user name
	if user.CheckUserName(userNew.UserName) == false {
		return "新用户登录名不符合命名规则", fmt.Errorf("UpdateUserPreCheck failed: new username  <%s> illegal", userNew.UserName)
	}

	//check exist
	userOld, ok := udm.userInfoByID[userNew.UserID]
	if ok == false {
		return "被更新的用户不存在", fmt.Errorf("UpdateUserPreCheck failed since user (userId = %d) not found", userNew.UserID)
	}

	//check duplicate name
	if userNew.UserName != userOld.UserName {
		if udm.IsUserNameExist(userNew.UserName) {
			return "新用户登录名与其他用户重名", fmt.Errorf("UpdateUserPreCheck failed since userName <%s> is already exist", userNew.UserName)
		}
	}

	if userOld.Role == user.RoleApprover {
		if user.CheckDepartmentCode(userNew.DepartmentCode) == false {
			return "新用户（审批人）设置的单位代码不符合命名规则",
				fmt.Errorf("UpdateUserPreCheck failed: user  DepartmentCode <%s> illegal", userNew.DepartmentCode)
		}

		if userOld.DepartmentCode != userNew.DepartmentCode {
			for _, u := range udm.userInfoByID {
				if u.DepartmentCode == userNew.DepartmentCode {
					return "新用户（审批人）设置的单位代码与其他单位的单位代码重复",
						fmt.Errorf("UpdateUserPreCheck failed: user  DepartmentCode <%s> duplicate", userNew.DepartmentCode)
				}
			}
		}

		if userOld.Department != userNew.Department {
			for _, u := range udm.userInfoByID {
				if u.Department == userNew.Department {
					return "新用户（审批人）设置的单位名称与其他单位的单位名称重复",
						fmt.Errorf("UpdateUserPreCheck failed: user  Department <%s> duplicate", userNew.Department)
				}
			}
		}
	}

	return "", nil
}

func (udm MemoryMap) UpdateUser(userNew user.UserInfo) (err error) {
	// check user name
	if user.CheckUserName(userNew.UserName) == false {
		return fmt.Errorf("UpdateUser failed: new username  <%s> illegal", userNew.UserName)
	}

	//check exist
	userOld, ok := udm.userInfoByID[userNew.UserID]
	if ok == false {
		return fmt.Errorf("UpdateUser failed since user (userId = %d) not found", userNew.UserID)
	}

	//update user info
	//general, only 6 property could be update
	err = udm.theUserDB.UpdateUser(userNew)
	if err != nil {
		return fmt.Errorf("udm.theUserDB.UpdateUser error: %v", err)
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
			err = udm.theUserDB.UpdateUserDepartment(userOld.DepartmentCode,
				userNew.Department, userNew.DepartmentCode)
			if err != nil {
				return fmt.Errorf("udm.theUserDB.UpdateUserDepartment error: %v", err)
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

func (udm MemoryMap) InsertUserPreCheck(userNew user.UserInfo) (string, error) {
	// Check username
	if user.CheckUserName(userNew.UserName) == false {
		return "新用户登录名不符合命名规则", fmt.Errorf("InsertUserPreCheck failed since userName  <%s> illegal", userNew.UserName)
	}

	// Check duplicate name
	if udm.IsUserNameExist(userNew.UserName) {
		return "新用户登录名与其他用户重名", fmt.Errorf("InsertUserPreCheck failed since userName <%s> is already exist", userNew.UserName)
	}

	//if user is Approver check department exist
	if userNew.Role == user.RoleApprover {
		if user.CheckDepartmentCode(userNew.DepartmentCode) == false {
			return "新用户（审批人）设置的单位代码不符合命名规则",
				fmt.Errorf("InsertUserPreCheck failed since user  DepartmentCode <%s> illegal", userNew.DepartmentCode)
		}

		for _, u := range udm.userInfoByName {
			if u.Department == userNew.Department || u.DepartmentCode == u.DepartmentCode {
				return "新用户（审批人）设置的单位名称与其他单位的单位名称重复",
					fmt.Errorf("InsertUserPreCheck failed since new Approver Department is already exist")
			}
		}
	}
	return "", nil
}

func (udm MemoryMap) InsertUser(userNew user.UserInfo) (err error) {
	//Insert to DB
	err = udm.theUserDB.InsertUser(userNew)
	if err != nil {
		return fmt.Errorf("udm.theUserDB.InsertUser error: %v", err)
	}

	//Query user from DB
	quser, err := udm.QueryUserByName(userNew.UserName)
	if err != nil {
		return fmt.Errorf("udm.QueryUserByName error: %v", err)
	}

	udm.userInfoByName[quser.UserName] = &quser
	udm.userInfoByID[quser.UserID] = &quser

	return nil
}

func (udm MemoryMap) DeleteUser(userID int) (err error) {
	userInfo, ok := udm.userInfoByID[userID]
	if !ok {
		return fmt.Errorf("user <%d> not exist", userID)
	}
	userName := userInfo.UserName

	// delete user from database
	err = udm.theUserDB.DeleteUser(userID)
	if err != nil {
		return fmt.Errorf("udm.theUserDB.DeleteUser error: %v", err)
	}

	// delete user from memory map
	udm.userInfoByID[userID] = nil
	udm.userInfoByName[userName] = nil

	return nil
}
