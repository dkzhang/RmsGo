package logMap

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var theLogMap = map[string]func() *logrus.Logger{
	"normal": getNormalLog,
}

func GetLog(name string) *logrus.Logger {
	if f, ok := theLogMap[name]; ok {
		return f()
	} else {
		return getDefaultLog(name)
	}
}

var normalLog *logrus.Logger
var normalOnce sync.Once

func getNormalLog() *logrus.Logger {
	normalOnce.Do(func() {
		fileName := "normal.log"
		if v, ok := theLogFileConfig.LogFile["normal"]; ok {
			fileName = v
		}
		//写入文件
		src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("err", err)
		}

		//实例化
		logger := logrus.New()

		//设置输出
		logger.Out = src

		//设置日志级别
		logger.SetLevel(logrus.DebugLevel)

		//设置日志格式
		logger.SetFormatter(&logrus.TextFormatter{})

	})
	return normalLog
}

func getDefaultLog(name string) *logrus.Logger {
	return nil
}
