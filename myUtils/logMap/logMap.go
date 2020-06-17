package logMap

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

//func GetLogArray(types ...string) []*logrus.Logger {
//	theLoggers := make([]*logrus.Logger, len(types))
//
//	for i, t := range types {
//		if l, ok := theLogMap[t]; ok {
//			theLoggers[i] = l
//		} else {
//			theLogMap[DEFAULT].WithFields(logrus.Fields{
//				"log name": t,
//			}).Info("Request a log name that has not been configured and returns the default GetLog.")
//			theLoggers[i] = theLogMap[DEFAULT]
//		}
//	}
//	return theLoggers
//}

const (
	NORMAL  = "normal"
	LOGIN   = "login"
	GIN     = "gin"
	DEFAULT = "default"
)

type LogMap struct {
	LogPathMap map[string]string           `yaml:"logPath,omitempty"`
	LoggerMap  map[string](*logrus.Logger) `yaml:"-"`
}

func NewLogMap(configFile string) (theLogMap LogMap) {
	var err error
	theLogMap, err = LoadLogConfig(configFile)
	if err != nil {
		logrus.Fatalf("LoadLogConfig error: %v", err)
	}
	logrus.Infof("NewLogMap LoadLogConfig success, LogPathMap = %v", theLogMap.LogPathMap)

	theLogMap.LoggerMap = make(map[string](*logrus.Logger), len(theLogMap.LogPathMap)+1)
	for t, p := range theLogMap.LogPathMap {
		theLogMap.LoggerMap[t] = loggerToFile(p, t+".log")
	}
	theLogMap.LoggerMap[DEFAULT] = logrus.New()

	logrus.Infof("NewLogMap make LoggerMap success, LoggerMap = %v", theLogMap.LoggerMap)

	return theLogMap
}

func (theLogMap LogMap) getLog(name string) *logrus.Logger {
	if v, ok := theLogMap.LoggerMap[name]; ok {
		return v
	} else {
		l := loggerToFile("./LogUnknownType"+name, name+".log")
		theLogMap.LoggerMap[name] = l
		return l
	}
}

func loggerToFile(logFilePath, logFileName string) *logrus.Logger {
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)

	////////////////////
	// 此处教程有误不应该打开文件，因为后续程序会以fileName作为文件名建立system link，
	// 如果这里创建打开了同名文件，会提示访问被拒。
	// 详细用法有待进一步研究
	// win10 下应该以管理员方式运行，否则提示没有权限创建链接
	// 写入文件
	//src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	//if err != nil {
	//	fmt.Println("err", err)
	//}
	////////////////////

	// 实例化
	logger := logrus.New()

	// 设置输出
	logger.Out = os.Stdout

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
	if err != nil {
		logrus.Fatalf("rotatelogs.New error: %v", err)
	}

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

func loggerToFile2(logFileName string) *logrus.Logger {
	// 实例化
	logger := logrus.New()

	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	// 设置 rotatelogs
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		logFileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(logFileName),

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

	logger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}

	logger.SetOutput(logWriter)

	return logger
}
