package allConfig

import (
	"fmt"
	"github.com/dkzhang/RmsGo/datebaseCommon/config"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func LoadAllConfig() (err error) {
	logMap.LoadLogConfig(os.Getenv("LogMapConf"))
	logMap.GetLog(logMap.DEFAULT).WithFields(logrus.Fields{
		"ENV LogMapConf": os.Getenv("LogMapConf"),
		"error":          err,
	}).Info("logMap.LoadLogConfig error.")

	theDbConfig, err := config.LoadDbConfig(os.Getenv("dbconf"))
	if err != nil {
		return fmt.Errorf("config.LoadDbConfig error: %v", err)
	}
	log.Printf("%v", theDbConfig)
	return nil
}
