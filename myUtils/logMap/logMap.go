package logMap

import (
	"github.com/sirupsen/logrus"
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
		normalLog = logrus.New()
	})
	return normalLog
}

func getDefaultLog(name string) *logrus.Logger {
	return nil
}
