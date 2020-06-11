package logMap

import "testing"

func TestLoggerToFile2(t *testing.T) {
	theLogger := loggerToFile2("C:\\Users\\52765\\go\\src\\github.com\\dkzhang\\RmsGo\\LogHere\\login.log")
	theLogger.Info(123)
	theLogger.Info("abc")
	theLogger.Info([]string{"a", "b"})
}

func TestLoggerToFile(t *testing.T) {
	theLogger := loggerToFile("C:\\Users\\dkzhang\\go\\src\\github.com\\dkzhang\\RmsGo\\LogHere", "login.log")
	theLogger.Info(123)
	theLogger.Info("abc")
	theLogger.Info([]string{"a", "b"})
}
