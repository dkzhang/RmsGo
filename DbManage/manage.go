package main

import (
	"fmt"
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreOps"
	databaseSecurity "github.com/dkzhang/RmsGo/datebaseCommon/security"
	"github.com/dkzhang/RmsGo/dbManage/pgManage"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("parameter input error. Expected at least 1 patameter.")
		return
	}

	/////////////////////////////////////////////////////////
	// Database: PostgreSQL
	theDbSecurity, err := databaseSecurity.LoadDbSecurity(os.Getenv("DbSE"))
	if err != nil {
		logMap.Log(logMap.DEFAULT).WithFields(logrus.Fields{
			"ENV DbSE": os.Getenv("DbConf"),
			"error":    err,
		}).Fatal("dbConfig.LoadDbSecurity error.")
		return
	}

	db, err := postgreOps.ConnectToDatabase(theDbSecurity.ThePgSecurity)
	if err != nil {
		logMap.Log(logMap.DEFAULT).WithFields(logrus.Fields{
			"error": err,
		}).Info("postgreSQL.ConnectToDatabase error.")
	} else {
		logMap.Log(logMap.DEFAULT).WithFields(logrus.Fields{
			"db": db,
		}).Info("postgreSQL.ConnectToDatabase success.")
	}
	defer db.Close()

	/////////////////////////////////////////////////////////

	switch os.Args[1] {
	case "create_all":
		fmt.Printf("删除所有表格并重建 \n")
		os.Setenv("DbConf", "./../Configuration/Security/database.yaml")
		PgManage.CreateAllTable(db)
	case "seed_all":
		fmt.Printf("无参：用测试数据初始化所有数据库表")
		PgManage.CreateAllTable(db)
		PgManage.SeedAllTable(db)
	case "import_from_file":
		fmt.Printf("表名，文件名：读取指定csv文件并将数据导入至指定数据表")
	}
}
