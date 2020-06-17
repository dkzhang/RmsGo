package userDB

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserInPostgre struct {
	db *sqlx.DB
}

func NewUserInPostgre(sqlxdb *sqlx.DB) (uip UserInPostgre) {
	return UserInPostgre{
		db: sqlxdb,
	}
}

func (upg UserInPostgre) Close() {
	upg.db.Close()
}

func (upg UserInPostgre) QueryUserByName(username string) (userInfo user.UserInfo, err error) {
	err = upg.db.Get(&userInfo, "SELECT * FROM user_info WHERE user_name=$1", username)
	if err != nil {
		return user.UserInfo{}, fmt.Errorf("query user in db error: %v", err)
	}
	return userInfo, nil
}

func (upg UserInPostgre) QueryUserByID(userID int) (userInfo user.UserInfo, err error) {
	err = upg.db.Get(&userInfo, "SELECT * FROM user_info WHERE user_id=$1", userID)
	if err != nil {
		return user.UserInfo{}, fmt.Errorf("query userInfo in db error: %v", err)
	}
	return userInfo, nil
}

func (upg UserInPostgre) GetAllUserInfo() (users []user.UserInfo, err error) {
	users = []user.UserInfo{}
	err = upg.db.Select(&users, "SELECT * FROM user_info")
	if err != nil {
		return nil, fmt.Errorf("get all user info from db error: %v", err)
	}
	return users, nil
}

func (upg UserInPostgre) UpdateUser(ui user.UserInfo) (err error) {
	//update user info
	//user_id, department, department_code and role update is not allowed.
	_, err = upg.db.NamedExec("UPDATE user_info "+
		"SET user_name=:user_name, chinese_name=:chinese_name, "+
		"section=:section, mobile=:mobile, "+
		"status=:status, remarks=:remarks "+
		"WHERE user_id=:user_id", ui)
	if err != nil {
		return fmt.Errorf("db.NamedExec UPDATE user_info error: %v", err)
	}
	return nil
}

func (upg UserInPostgre) UpdateUserDepartment(oldDepCode string, newDep string, newDepCode string) (err error) {
	//update user department with same department_code
	_, err = upg.db.NamedExec("UPDATE user_info "+
		"SET department=:department, department_code=:department_code "+
		"WHERE department_code=:old_dep_code",
		map[string]interface{}{
			"department":      newDep,
			"department_code": newDepCode,
			"old_dep_code":    oldDepCode,
		})
	if err != nil {
		return fmt.Errorf("db.NamedExec UPDATE user_info error: %v", err)
	}
	return nil
}

func (upg UserInPostgre) InsertUser(ui user.UserInfo) (err error) {
	insertUser := `INSERT INTO user_info (user_name, chinese_name, department, department_code, section, mobile, role, status,remarks) VALUES (:user_name, :chinese_name, :department, :department_code, :section, :mobile, :role, :status, :remarks)`
	result, err := upg.db.NamedExec(insertUser, ui)
	if err != nil {
		return fmt.Errorf("db.NamedExec(insertUser, ui), UserName = %s, UserInfo = %v :%v", ui.UserName, ui, err)
	}
	fmt.Printf("InsertUser success: %v \n", result)
	return nil
}

func (upg UserInPostgre) DeleteUser(userID int) (err error) {
	deleteUser := `DELETE FROM user_info WHERE user_id=$1`

	result, err := upg.db.Exec(deleteUser, userID)
	if err != nil {
		return fmt.Errorf("db.Exec(deleteUser, userID), userID = %d", userID)
	}
	fmt.Printf("DeleteUser success: %v \n", result)
	return nil
}
