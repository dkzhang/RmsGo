package allConfig

import (
	dbConfig "github.com/dkzhang/RmsGo/datebaseCommon/config"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/sirupsen/logrus"
	"os"
)

// 读取所有日志配置
func LoadAllConfig() {
	var err error

	err = logMap.LoadLogConfig(os.Getenv("LogMapConf"))
	if err != nil {
		logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
			"ENV LogMapConf": os.Getenv("LogMapConf"),
			"error":          err,
		}).Error("logMap.LoadLogConfig error.")
	}

	err = dbConfig.LoadDbConfig(os.Getenv("DbConf"))
	if err != nil {
		logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
			"ENV DbConf": os.Getenv("DbConf"),
			"error":      err,
		}).Fatal("dbConfig.LoadDbConfig error.")
		return
	}

	return
}
