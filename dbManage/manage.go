package main

import (
	"fmt"
	"github.com/dkzhang/RmsGo/allConfig"
	"github.com/dkzhang/RmsGo/datebaseCommon/config"
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreSQL"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("parameter input error. Expected at least 1 patameter.")
		return
	}

	switch os.Args[1] {
	case "create_all":
		fmt.Printf("删除所有表格并重建 \n")
		createAll()
	case "seed_all":
		fmt.Printf("无参：用测试数据初始化所有数据库表")
	case "import_from_file":
		fmt.Printf("表名，文件名：读取指定csv文件并将数据导入至指定数据表")
	}
}

func createAll() {
	os.Setenv("DbConf", "./../Configuration/Security/database.yaml")

	allConfig.LoadAllConfig()
	logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
		"config.TheDbConfig": config.TheDbConfig,
	}).Info("allConfig.LoadAllConfig success.")

	db, err := postgreSQL.ConnectToDatabase(config.TheDbConfig.ThePgConfig)
	if err != nil {
		logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
			"error": err,
		}).Info("postgreSQL.ConnectToDatabase error.")
	} else {
		logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
			"db": db,
		}).Info("postgreSQL.ConnectToDatabase success.")
	}
	defer db.Close()
}
