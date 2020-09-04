package myPgBytea

import (
	"fmt"
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreOpsSqlx"
	databaseSecurity "github.com/dkzhang/RmsGo/datebaseCommon/security"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var SchemaUser = `
		CREATE TABLE user_info_dktest (
    		user_id SERIAL PRIMARY KEY,
			user_name bytea,				
			remarks varchar(256)
		);
		`

type UserInfo struct {
	UserID   int    `db:"user_id" json:"user_id"`
	UserName []byte `db:"user_name" json:"user_name"`
	Remarks  string `db:"remarks" json:"remarks"`
}

func ConnectToDatabase(conf string) *sqlx.DB {
	theDbSecurity, err := databaseSecurity.LoadDbSecurity(conf)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"conf":  conf,
			"error": err,
		}).Fatal("dbConfig.LoadDbSecurity error.")
	}

	theDb, err := postgreOpsSqlx.ConnectToDatabase(theDbSecurity.ThePgSecurity)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ThePgSecurity": theDbSecurity.ThePgSecurity,
			"error":         err,
		}).Fatal("postgreOps.ConnectToDatabase error.")
	}
	return theDb
}

func CreateTable(db *sqlx.DB) {
	_, err := db.Exec(SchemaUser)
	if err != nil {
		logrus.Fatal("Create table error")
	}
}

func DropTable(db *sqlx.DB) {
	_, err := db.Exec(`DROP Table user_info_dktest`)
	if err != nil {
		logrus.Fatal("Create table error")
	}
}

func InsertUser(db *sqlx.DB, theUser UserInfo) (err error) {
	insertUser := `INSERT INTO user_info_dktest (user_name, remarks) VALUES (:user_name, :remarks)`
	result, err := db.NamedExec(insertUser, theUser)
	if err != nil {
		return fmt.Errorf("db.NamedExec(insertUser, theUser) error, UserInfo = %v :%v", theUser, err)
	}
	fmt.Printf("InsertUser success: %v \n", result)
	return nil
}

func RetrieveUser(db *sqlx.DB, userID int) (theUser UserInfo, err error) {
	err = db.Get(&theUser, "SELECT * FROM user_info_dktest WHERE user_id=$1", userID)
	if err != nil {
		return UserInfo{}, fmt.Errorf("query userInfo in db error: %v", err)
	}
	return theUser, nil
}
