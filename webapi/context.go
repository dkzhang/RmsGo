package webapi

import (
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreOps"
	"github.com/dkzhang/RmsGo/datebaseCommon/redisOps"
	databaseSecurity "github.com/dkzhang/RmsGo/datebaseCommon/security"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	userConfig "github.com/dkzhang/RmsGo/webapi/dataManagement/userTempDM/config"
	userSecurity "github.com/dkzhang/RmsGo/webapi/dataManagement/userTempDM/security"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var TheContext Context

type Context struct {
	TheDbConfig databaseSecurity.DbSecurity
	TheDb       *sqlx.DB
	TheRedis    *redisOps.Redis

	TheLoginConfig   userConfig.LoginConfig
	TheLoginSecurity userSecurity.LoginSecurity
}

var initOnce sync.Once

func InitInfrastructure() {
	initOnce.Do(func() {
		var err error

		TheContext.TheDbConfig, err = databaseSecurity.LoadDbSecurity(os.Getenv("DbConf"))
		if err != nil {
			logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
				"ENV DbConf": os.Getenv("DbConf"),
				"error":      err,
			}).Fatal("dbConfig.LoadDbSecurity error.")
			return
		}

		TheContext.TheDb, err = postgreOps.ConnectToDatabase(TheContext.TheDbConfig.ThePgConfig)
		if err != nil {
			logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
				"ThePgConfig": TheContext.TheDbConfig.ThePgConfig,
				"error":       err,
			}).Fatal("postgreOps.ConnectToDatabase error.")
			return
		}

		TheContext.TheLoginConfig, err = userConfig.LoadLoginConfig(os.Getenv("LoginConf"))
		if err != nil {
			logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
				"ENV LoginConf": os.Getenv("LogMapConf"),
				"error":         err,
			}).Fatalf("userConfig.LoadLoginSecurity error.")
		}

		TheContext.TheLoginSecurity, err = userSecurity.LoadLoginSecurity()
		if err != nil {
			logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
				"error": err,
			}).Fatalf("userConfig.LoadLoginSecurity error.")
		}
	})
}
