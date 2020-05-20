package user

import (
	"database/sql"
	"fmt"
	"log"

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
