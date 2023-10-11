package log_test

import (
	"net/http"
	"testing"

	"github.com/guru-invest/guru.framework/src/helpers/log"
)

func createLoggerWithHeaders() *log.GuruLog {
	headers := http.Header{}
	headers.Add("user-agent", "Version | teste2 |")
	headers.Add("device-id", "ASD4AS56D4-4ASD54AS6D-4AS5D4AS65")
	headers.Add("X-Forwarded-For", "192.168.0.1")
	headers.Add("service-name", "guru.framework")
	headers.Add("correlation-id", "AAAAA1-AAAAA2-AAAAA3")
	headers.Add("session-id", "FD45D84F7E")

	return log.NewGuruLog("unit teste", log.LevelInfo, log.WithHTTPHeader(headers))
}

func TestLogDefaultInfoFieldsNull(t *testing.T) {
	guruLogger := log.NewGuruLog("unit teste", log.LevelInfo)
	guruLogger.Info("primeiro log", nil)

	t.Log("OK")
}

func TestLogInfo(t *testing.T) {
	guruLogger := createLoggerWithHeaders()
	fields := log.Fields{
		"CustomerCode": "customerCode",
		"Caller":       "TestLogInfo",
		"msg 1":        "mensagem 1",
		"msg 2":        "mensagem 2",
		"msg 3":        "mensagem 3",
	}
	guruLogger.Info("primeiro log", fields)
	t.Log("OK")
}

func TestLogDebug(t *testing.T) {
	guruLogger := createLoggerWithHeaders()
	fields := log.Fields{
		"CustomerCode": "customerCode",
		"Caller":       "TestLogDebug",
		"msg 1":        "mensagem 1",
		"msg 2":        "mensagem 2",
		"msg 3":        "mensagem 3",
	}
	guruLogger.Debug("primeiro log", fields)
	t.Log("OK")
}

func TestLogError(t *testing.T) {
	guruLogger := createLoggerWithHeaders()
	fields := log.Fields{
		"CustomerCode": "customerCode",
		"Caller":       "TestLogError",
		"msg 1":        "mensagem 1",
		"msg 2":        "mensagem 2",
		"msg 3":        "mensagem 3",
	}
	guruLogger.Error("primeiro log", fields)
	t.Log("OK")
}
