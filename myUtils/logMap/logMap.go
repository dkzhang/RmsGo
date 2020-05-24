package logMap

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"sync"
	"time"
)

const NORMAL = "normal"
const DEFAULT = "default"

var theLogMap = map[string]func() *logrus.Logger{
	NORMAL:  getNormalLog,
	DEFAULT: getDefaultLog,
}

var defaultLog = logrus.New()

func GetLog(name string) *logrus.Logger {
	if l, ok := theLogMap[name]; ok {
		return l()
	} else {
		defaultLog.WithFields(logrus.Fields{
			"log name": name,
		}).Info("Request a log name that has not been configured and returns the default Log.")
		return getDefaultLog()
	}
}

var normalLog *logrus.Logger
var normalOnce sync.Once

func getNormalLog() *logrus.Logger {
	normalOnce.Do(func() {
		filePath := ""
		fileName := NORMAL + ".log"
		if v, ok := theLogFileConfig.LogFile[NORMAL]; ok {
			filePath = v
		}
		normalLog = loggerToFile(filePath, fileName)

	})
	return normalLog
}

func getDefaultLog() *logrus.Logger {
	return defaultLog
}

func loggerToFile(logFilePath, logFileName string) *logrus.Logger {
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)

	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	// 实例化
	logger := logrus.New()

	// 设置输出
	logger.Out = src

	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),

		//////////////////////////////////////////////
		// WithLinkName为最新的日志建立软连接,以方便随着找到当前日志文件
		// WithRotationTime设置日志分割的时间,这里设置为一小时分割一次
		// WithMaxAge和WithRotationCount二者只能设置一个
		// WithMaxAge设置文件清理前的最长保存时间
		// WithRotationCount设置文件清理前最多保存的个数.

	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	logger.AddHook(lfHook)

	return logger
}
