package databaseInit

import (
	"fmt"
	"github.com/dkzhang/RmsGo/databaseInit/pgOps"
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreOps"
	databaseSecurity "github.com/dkzhang/RmsGo/datebaseCommon/security"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"os"
)

func CreateAllTable() {
	fmt.Printf("删除所有表格并重建 \n")
	db := connectToDatabase()
	pgOps.CreateAllTable(db)
}

func SeedAllTable() {
	fmt.Printf("用测试数据初始化所有数据库表")
	db := connectToDatabase()
	pgOps.CreateAllTable(db)
	pgOps.SeedAllTable(db)
}

func ImportFromFile(tableName, fileName string) {
	fmt.Printf("表名，文件名：读取指定csv文件并将数据导入至指定数据表")
	db := connectToDatabase()
	pgOps.CreateAllTable(db)
	pgOps.ImportFromFile(db)
}

func connectToDatabase() (db *sqlx.DB) {
	theDbSecurity, err := databaseSecurity.LoadDbSecurity(os.Getenv("DbSE"))
	if err != nil {
		logMap.Log(logMap.DEFAULT).WithFields(logrus.Fields{
			"ENV DbSE": os.Getenv("DbSE"),
			"error":    err,
		}).Fatal("dbConfig.LoadDbSecurity error.")
		return
	}

	db, err = postgreOps.ConnectToDatabase(theDbSecurity.ThePgSecurity)
	if err != nil {
		logMap.Log(logMap.DEFAULT).WithFields(logrus.Fields{
			"error": err,
		}).Fatal("postgreSQL.ConnectToDatabase error.")
	} else {
		logMap.Log(logMap.DEFAULT).WithFields(logrus.Fields{
			"db": db,
		}).Info("postgreSQL.ConnectToDatabase success.")
	}
	return db
}
