package webapi

import (
	"github.com/dkzhang/RmsGo/datebaseCommon/redisOps"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/dataManagement/userTempDM"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var TheContext Context

type Context struct {
	TheDb    *sqlx.DB
	TheRedis *redisOps.Redis

	TheLoginConfig   userTempDM.LoginConfig
	TheLoginSecurity userTempDM.LoginSecurity
}

var initOnce sync.Once

func InitInfrastructure() {
	initOnce.Do(func() {
		var err error

		TheContext.TheLoginConfig, err = userTempDM.LoadLoginConfig(os.Getenv("LoginConf"))
		if err != nil {
			logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
				"ENV LoginConf": os.Getenv("LogMapConf"),
				"error":         err,
			}).Fatalf("userTempDM.LoadLoginSecurity error.")
		}

		TheContext.TheLoginSecurity, err = userTempDM.LoadLoginSecurity()
		if err != nil {
			logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
				"error": err,
			}).Fatalf("userTempDM.LoadLoginSecurity error.")
		}
	})
}
