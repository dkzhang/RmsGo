package PgManage

import (
	"database/sql"
	"github.com/dkzhang/RmsGo/allConfig"
	"github.com/dkzhang/RmsGo/datebaseCommon/config"
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreOps"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var tableList = map[string]string{
	"user_info": user.SchemaUser,
}

func createTable(db *sqlx.DB, schema string) (result sql.Result, err error) {
	result, err = db.Exec(schema)
	return result, err
}

func dropTable(db *sqlx.DB, tableName string) (result sql.Result, err error) {
	exec := `DROP Table ` + tableName
	result, err = db.Exec(exec)
	return result, err
}

func CreateAllTable() {
	allConfig.LoadAllConfig()
	logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
		"config.TheDbConfig": security.TheDbConfig,
	}).Info("allConfig.LoadAllConfig success.")

	db, err := postgreOps.ConnectToDatabase(security.TheDbConfig.ThePgConfig)
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

	for name, scheme := range tableList {
		dropTable(db, name)
		createTable(db, scheme)
	}
}
