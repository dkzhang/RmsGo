package infrastructure

import (
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreOpsSqlx"
	"github.com/dkzhang/RmsGo/datebaseCommon/redisOps"
	databaseSecurity "github.com/dkzhang/RmsGo/datebaseCommon/security"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/myUtils/shortMessageService"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/applicationDB"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/applicationDM"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/generalFormDraftDB"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDB"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/projectDM"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/userDB"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/userDM"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/userTempDM"
	userConfig "github.com/dkzhang/RmsGo/webapi/dataInfra/userTempDM/config"
	userSecurity "github.com/dkzhang/RmsGo/webapi/dataInfra/userTempDM/security"
	"github.com/dkzhang/RmsGo/webapi/handle/extractLoginUserInfo"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/dkzhang/RmsGo/webapi/model/project"
	"github.com/dkzhang/RmsGo/webapi/workflow"
	"github.com/dkzhang/RmsGo/webapi/workflow/applyProjectAndResource"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

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

	TheExtractor extractLoginUserInfo.Extractor

	TheGeneralFormDraftDB generalFormDraftDB.GeneralFormDraftDB

	TheApplicationDB applicationDB.ApplicationDB
	TheApplicationDM applicationDM.ApplicationDM

	TheProjectDB projectDB.ProjectDB
	TheProjectDM projectDM.ProjectDM

	TheLogMap logMap.LogMap

	TheApplyProjectAndResourceWorkflow workflow.GeneralWorkflow
}

type InfraConfigFile struct {
	LogMapConf string
	SmsSE      string
	DbSE       string
	LoginConf  string
}

func NewInfrastructure(icf InfraConfigFile) *Infrastructure {
	theInfras := Infrastructure{}

	var err error

	/////////////////////////////////////////////////////////
	// LOG
	theInfras.TheLogMap = logMap.NewLogMap(icf.LogMapConf)

	/////////////////////////////////////////////////////////
	// SMS
	theSmsSecurity, err := shortMessageService.LoadSmsSecurity(icf.SmsSE)
	if err != nil {
		theInfras.TheLogMap.Log(logMap.DEFAULT).WithFields(logrus.Fields{
			"ENV SmsSE": icf.SmsSE,
			"error":     err,
		}).Fatal("shortMessageService.LoadSmsSecurity error.")
	}
	theInfras.TheSmsService = shortMessageService.NewSmsTencentCloudService(theSmsSecurity)

	/////////////////////////////////////////////////////////
	// Database: PostgreSQL and Redis
	theInfras.TheDbSecurity, err = databaseSecurity.LoadDbSecurity(icf.DbSE)
	if err != nil {
		theInfras.TheLogMap.Log(logMap.DEFAULT).WithFields(logrus.Fields{
			"ENV DbSE": icf.DbSE,
			"error":    err,
		}).Fatal("dbConfig.LoadDbSecurity error.")
	}

	theInfras.TheDb, err = postgreOpsSqlx.ConnectToDatabase(theInfras.TheDbSecurity.ThePgSecurity)
	if err != nil {
		theInfras.TheLogMap.Log(logMap.DEFAULT).WithFields(logrus.Fields{
			"ThePgSecurity": theInfras.TheDbSecurity.ThePgSecurity,
			"error":         err,
		}).Fatal("postgreOps.ConnectToDatabase error.")
	}

	opts := &redisOps.RedisOpts{
		Host: theInfras.TheDbSecurity.TheRedisSecurity.Host,
	}
	theInfras.TheRedis = redisOps.NewRedis(opts)

	/////////////////////////////////////////////////////////
	// Login and UserTempDM
	theInfras.TheLoginConfig, err = userConfig.LoadLoginConfig(icf.LoginConf)
	if err != nil {
		theInfras.TheLogMap.Log(logMap.DEFAULT).WithFields(logrus.Fields{
			"ENV LoginConf": icf.LoginConf,
			"error":         err,
		}).Fatal("userConfig.LoadLoginSecurity error.")
	}

	theInfras.TheLoginSecurity, err = userSecurity.LoadLoginSecurity()
	if err != nil {
		theInfras.TheLogMap.Log(logMap.DEFAULT).WithFields(logrus.Fields{
			"error": err,
		}).Fatal("userConfig.LoadLoginSecurity error.")
	}

	theInfras.TheUserTempDM = userTempDM.NewRedisAndJwt(theInfras.TheRedis,
		theInfras.TheLoginConfig, theInfras.TheLoginSecurity)

	/////////////////////////////////////////////////////////
	// UserDM and UserDB
	theInfras.TheUserDB = userDB.NewUserInPostgre(theInfras.TheDb)
	theInfras.TheUserDM, err = userDM.NewMemoryMap(theInfras.TheUserDB, theInfras.TheLogMap)
	if err != nil {
		theInfras.TheLogMap.Log(logMap.DEFAULT).WithFields(logrus.Fields{
			"error": err,
		}).Fatal("userDM.NewMemoryMap error.")
	}

	/////////////////////////////////////////////////////////
	// GeneralFormDraftDB
	theInfras.TheGeneralFormDraftDB = generalFormDraftDB.NewGeneralFormDraftPg(theInfras.TheDb)

	/////////////////////////////////////////////////////////
	// Application DB & DM
	theInfras.TheApplicationDB = applicationDB.NewApplicationPg(theInfras.TheDb, application.TableApp, application.TableAppOps)
	theInfras.TheApplicationDM, err = applicationDM.NewMemoryMap(theInfras.TheApplicationDB, theInfras.TheLogMap)
	if err != nil {
		theInfras.TheLogMap.Log(logMap.DEFAULT).WithFields(logrus.Fields{
			"error": err,
		}).Fatal("applicationDM.NewMemoryMap error.")
	}

	/////////////////////////////////////////////////////////
	// Project DB & DM
	theInfras.TheProjectDB = projectDB.NewProjectPg(theInfras.TheDb, project.TableName)
	theInfras.TheProjectDM, err = projectDM.NewMemoryMap(theInfras.TheProjectDB, theInfras.TheLogMap)
	if err != nil {
		theInfras.TheLogMap.Log(logMap.DEFAULT).WithFields(logrus.Fields{
			"error": err,
		}).Fatal("projectDM.NewMemoryMap error.")
	}

	/////////////////////////////////////////////////////////
	// Extractor
	theInfras.TheExtractor = extractLoginUserInfo.NewExtractor(theInfras.TheLogMap, theInfras.TheUserDM)

	/////////////////////////////////////////////////////////
	// Workflow
	theInfras.TheApplyProjectAndResourceWorkflow = ApplyProjectAndResource.NewWorkflow(theInfras.TheApplicationDM, theInfras.TheProjectDM)

	return &theInfras
}
