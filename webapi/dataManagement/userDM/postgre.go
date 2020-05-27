package userDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func QueryUserByName(username string, db *sqlx.DB) (userInfo user.UserInfo, err error) {
	err = db.Get(&userInfo, "SELECT * FROM user_info WHERE user_name=$1", username)
	if err != nil {
		return user.UserInfo{}, fmt.Errorf("query user in db error: %v", err)
	}
	return userInfo, nil
}

func QueryUserByID(userID int, db *sqlx.DB) (userInfo user.UserInfo, err error) {
	err = db.Get(&userInfo, "SELECT * FROM user_info WHERE user_id=$1", userID)
	if err != nil {
		return user.UserInfo{}, fmt.Errorf("query userInfo in db error: %v", err)
	}
	return userInfo, nil
}

func GetAllUserInfo(db *sqlx.DB) (users []user.UserInfo, err error) {
	users = []user.UserInfo{}
	err = db.Select(&users, "SELECT * FROM user_info")
	if err != nil {
		return nil, fmt.Errorf("get all user info from db error: %v", err)
	}
	return users, nil
}

func UpdateUser(ui user.UserInfo, db *sqlx.DB) (isNoDuplicateName bool, err error) {
	isNoDuplicateName, err = VerifyNoDuplicateName(ui, db)
	if isNoDuplicateName == false || err != nil {
		return isNoDuplicateName, err
	}

	//update user info
	//user_id, department, department_code and role update is not allowed.
	_, err = db.NamedExec("UPDATE user_info "+
		"SET user_name=:user_name, chinese_name=:chinese_name, "+
		"section=:section, mobile=:mobile, "+
		"status=:status, remarks=:remarks "+
		"WHERE user_id=:user_id", ui)
	if err != nil {
		return true, fmt.Errorf("db.NamedExec UPDATE user_info error: %v", err)
	}
	return true, nil
}

func UpdateUserDepartment(depCode string, newDep string, db *sqlx.DB) (err error) {
	//update user department with same department_code
	_, err = db.NamedExec("UPDATE user_info "+
		"SET department=:department "+
		"WHERE department_code=:department_code",
		map[string]interface{}{
			"department":      newDep,
			"department_code": depCode,
		})
	if err != nil {
		return fmt.Errorf("db.NamedExec UPDATE user_info error: %v", err)
	}
	return nil
}

func VerifyNoDuplicateName(ui user.UserInfo, db *sqlx.DB) (isNoDuplicateName bool, err error) {
	tempUser := user.UserInfo{}
	err = db.Get(&tempUser, "SELECT * FROM user_info WHERE user_name = $1", ui.UserName)
	if err == nil {
		//如果找到了有相同user_name的记录，则说明新记录重名了
		return false, fmt.Errorf("duplicate UserName: %s", ui.UserName)
	}
	return true, nil
}

func InsertUser(ui user.UserInfo, db *sqlx.DB) (isNoDuplicateName bool, err error) {
	isNoDuplicateName, err = VerifyNoDuplicateName(ui, db)
	if isNoDuplicateName == false || err != nil {
		return isNoDuplicateName, err
	}

	insertUser := `INSERT INTO user_info (user_name, chinese_name, department, department_code, section, mobile, role, status,remarks) VALUES (:user_name, :chinese_name, :department, :department_code, :section, :mobile, :role, :status, :remarks)`
	result, err := db.NamedExec(insertUser, ui)
	if err != nil {
		return true, fmt.Errorf("db.NamedExec(insertUser, ui), UserName = %s, UserInfo = %v :%v", ui.UserName, ui, err)
	}
	fmt.Printf("InsertUser success: %v", result)
	return true, nil
}
