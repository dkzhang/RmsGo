package logMap

import "github.com/sirupsen/logrus"

type DkLog struct {
	loggers []*logrus.Logger
	fields  logrus.Fields
}

func Log(types ...string) (dl DkLog) {
	dl = DkLog{
		loggers: make([]*logrus.Logger, len(types)),
		fields:  nil,
	}

	if len(types) == 0 {
		dl.loggers = []*logrus.Logger{theLogMap[DEFAULT]}
	} else {
		for i, t := range types {
			if l, ok := theLogMap[t]; ok {
				dl.loggers[i] = l
			} else {
				theLogMap[DEFAULT].WithFields(logrus.Fields{
					"log name": t,
				}).Info("Request a log name that has not been configured and returns the default GetLog.")
				dl.loggers[i] = theLogMap[DEFAULT]
			}
		}
	}

	return dl
}

func (dl DkLog) WithFields(fields logrus.Fields) DkLog {
	return DkLog{
		loggers: dl.loggers,
		fields:  fields,
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////

func (dl DkLog) Info(args ...interface{}) {
	if dl.fields != nil {
		for _, l := range dl.loggers {
			l.WithFields(dl.fields).Info(args...)
		}
	} else {
		for _, l := range dl.loggers {
			l.Info(args...)
		}
	}
}

func (dl DkLog) Error(args ...interface{}) {
	if dl.fields != nil {
		for _, l := range dl.loggers {
			l.WithFields(dl.fields).Error(args...)
		}
	} else {
		for _, l := range dl.loggers {
			l.Error(args...)
		}
	}
}

func (dl DkLog) Fatal(args ...interface{}) {
	if dl.fields != nil {
		for _, l := range dl.loggers {
			l.WithFields(dl.fields).Fatal(args...)
		}
	} else {
		for _, l := range dl.loggers {
			l.Fatal(args...)
		}
	}
}
