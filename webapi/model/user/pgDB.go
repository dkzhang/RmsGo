package user

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func QueryUser(username string, db *sqlx.DB) (user UserInfo, err error) {
	err = db.Get(&user, "SELECT * FROM user_info WHERE user_name=$1", username)
	if err != nil {
		return UserInfo{}, fmt.Errorf("query user in db error: %v", err)
	}
	return user, nil
}

func GetAllUserInfo(db *sqlx.DB) (users []UserInfo, err error) {
	users = []UserInfo{}
	err = db.Select(&users, "SELECT * FROM user_info")
	if err != nil {
		return nil, fmt.Errorf("get all user info from db error: %v", err)
	}
	return users, nil
}

func UpdateUser(ui UserInfo, db *sqlx.DB) (err error) {

	db.NamedExec("UPDATE user_info SET user_name=:UserName WHERE user_id=:UserID", ui)
	return nil
}

func InsertUser(ui UserInfo, db *sqlx.DB) (err error) {
	insertUser := `INSERT INTO user_info (user_name, chinese_name, department, department_code, section, mobile, role, status,remarks) VALUES (:user_name, :chinese_name, :department, :department_code, :section, :mobile, :role, :status, :remarks)`

	result, err := db.NamedExec(insertUser, ui)
	if err != nil {
		return fmt.Errorf("db.NamedExec(insertUser, ui), UserName = %s, UserInfo = %v :%v", ui.UserName, ui, err)
	}
	fmt.Printf("InsertUser success: %v", result)
	return nil
}
