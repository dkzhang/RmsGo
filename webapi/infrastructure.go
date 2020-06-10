package webapi

import (
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreOps"
	"github.com/dkzhang/RmsGo/datebaseCommon/redisOps"
	databaseSecurity "github.com/dkzhang/RmsGo/datebaseCommon/security"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/myUtils/shortMessageService"
	"github.com/dkzhang/RmsGo/webapi/dataManagement/userDB"
	"github.com/dkzhang/RmsGo/webapi/dataManagement/userDM"
	"github.com/dkzhang/RmsGo/webapi/dataManagement/userTempDM"
	userConfig "github.com/dkzhang/RmsGo/webapi/dataManagement/userTempDM/config"
	userSecurity "github.com/dkzhang/RmsGo/webapi/dataManagement/userTempDM/security"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var TheInfras Infrastructure

type Infrastructure struct {
	TheSmsService shortMessageService.SmsService

	TheDbSecurity databaseSecurity.DbSecurity
	TheDb         *sqlx.DB
	TheRedis      *redisOps.Redis

	TheLoginConfig   userConfig.LoginConfig
	TheLoginSecurity userSecurity.LoginSecurity

	TheUserDB     userDB.UserDB
	TheUserDM     userDM.UserDM
	TheUserTempDM userTempDM.UserTempDM
}

var initOnce sync.Once

func InitInfrastructure() {
	initOnce.Do(func() {
		var err error

		/////////////////////////////////////////////////////////
		// LOG
		err = logMap.LoadLogConfig(os.Getenv("LogMapConf"))
		if err != nil {
			logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
				"ENV LogMapConf": os.Getenv("LogMapConf"),
				"error":          err,
			}).Error("logMap.LoadLogConfig error.")
		}

		/////////////////////////////////////////////////////////
		// SMS
		theSmsSecurity, err := shortMessageService.LoadSmsSecurity(os.Getenv("SmsSE"))
		if err != nil {
			logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
				"ENV SmsSE": os.Getenv("SmsSE"),
				"error":     err,
			}).Error("shortMessageService.LoadSmsSecurity error.")
		}
		TheInfras.TheSmsService = shortMessageService.NewSmsService(theSmsSecurity)

		/////////////////////////////////////////////////////////
		// Database: PostgreSQL and Redis
		TheInfras.TheDbSecurity, err = databaseSecurity.LoadDbSecurity(os.Getenv("DbSE"))
		if err != nil {
			logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
				"ENV DbSE": os.Getenv("DbConf"),
				"error":    err,
			}).Fatal("dbConfig.LoadDbSecurity error.")
			return
		}

		TheInfras.TheDb, err = postgreOps.ConnectToDatabase(TheInfras.TheDbSecurity.ThePgSecurity)
		if err != nil {
			logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
				"ThePgSecurity": TheInfras.TheDbSecurity.ThePgSecurity,
				"error":         err,
			}).Fatal("postgreOps.ConnectToDatabase error.")
			return
		}

		opts := &redisOps.RedisOpts{
			Host: TheInfras.TheDbSecurity.TheRedisSecurity.Host,
		}
		TheInfras.TheRedis = redisOps.NewRedis(opts)

		/////////////////////////////////////////////////////////
		// Login and UserTempDM
		TheInfras.TheLoginConfig, err = userConfig.LoadLoginConfig(os.Getenv("LoginConf"))
		if err != nil {
			logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
				"ENV LoginConf": os.Getenv("LogMapConf"),
				"error":         err,
			}).Fatalf("userConfig.LoadLoginSecurity error.")
		}

		TheInfras.TheLoginSecurity, err = userSecurity.LoadLoginSecurity()
		if err != nil {
			logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
				"error": err,
			}).Fatalf("userConfig.LoadLoginSecurity error.")
		}

		TheInfras.TheUserTempDM = userTempDM.NewRedisAndJwt(TheInfras.TheRedis,
			TheInfras.TheLoginConfig, TheInfras.TheLoginSecurity)

		/////////////////////////////////////////////////////////
		// UserDM and UserDB
		TheInfras.TheUserDB = userDB.NewUserInPostgre(TheInfras.TheDb)
		TheInfras.TheUserDM, err = userDM.NewMemoryMap(TheInfras.TheUserDB)
		if err != nil {
			logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
				"error": err,
			}).Fatalf("userDM.NewMemoryMap error.")
		}

	})
}
