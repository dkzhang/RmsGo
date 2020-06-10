package logMap

import "testing"

func TestLoggerToFile(t *testing.T) {
	theLogger := LoggerToFile("C:\\Users\\52765\\go\\src\\github.com\\dkzhang\\RmsGo\\LogHere\\login.log")
	theLogger.Info(123)
	theLogger.Info("abc")
	theLogger.Info([]string{"a", "b"})
}
