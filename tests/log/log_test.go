package log_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/guru-invest/guru.framework/src/helpers/log"
)

func getGuruLog() log.GuruLog {
	guruLog := log.GuruLog{
		HTTPHeader: &http.Header{},
	}

	guruLog.HTTPHeader.Add("user-agent", "Version | teste2 |")
	guruLog.HTTPHeader.Add("device-id", "ASD4AS56D4-4ASD54AS6D-4AS5D4AS65")
	guruLog.HTTPHeader.Add("X-Forwarded-For", "192.168.0.1")
	guruLog.HTTPHeader.Add("service-name", "guru.framework")
	guruLog.HTTPHeader.Add("correlation-id", "AAAAA1-AAAAA2-AAAAA3")
	guruLog.HTTPHeader.Add("session-id", "FD45D84F7E")

	return guruLog
}

func TestLogDefaultInfoFieldsNull(t *testing.T) {

	log.InitLog("DEFAULT", "unit teste")

	guruEmptyLog := log.GuruLog{
		HTTPHeader: nil,
	}

	guruEmptyLog.Info(nil, "primeiro log")

	fmt.Println("OK")
}

func TestLogInfo(t *testing.T) {
	log.InitLog("INFO", "unit teste")

	guruLog := getGuruLog()

	guruLog.Info(&log.LogWithFields{
		CustomerCode: "customerCode",
		Caller:       "TestLogInfo",
		InfoMessage: log.Fields{
			"msg 1": "mensagem 1",
			"msg 2": "mensagem 2",
			"msg 3": "mensagem 3",
		},
	}, "primeiro log",
	)

	fmt.Println("OK")
}

func TestLogDebug(t *testing.T) {
	log.InitLog("DEBUG", "unit teste")

	guruLog := getGuruLog()

	guruLog.Debug(&log.LogWithFields{
		CustomerCode: "customerCode",
		Caller:       "TestLogDebug",
		InfoMessage: log.Fields{
			"msg 1": "mensagem 1",
			"msg 2": "mensagem 2",
			"msg 3": "mensagem 3",
		},
	}, "primeiro log",
	)

	fmt.Println("OK")
}

func TestLogError(t *testing.T) {
	log.InitLog("ERROR", "unit teste")

	guruLog := getGuruLog()

	guruLog.Error(&log.LogWithFields{
		CustomerCode: "customerCode",
		Caller:       "TestLogError",
		InfoMessage: log.Fields{
			"msg 1": "mensagem 1",
			"msg 2": "mensagem 2",
			"msg 3": "mensagem 3",
		},
	}, "primeiro log",
	)

	fmt.Println("OK")
}

func TestLogWarn(t *testing.T) {

	log.InitLog("WARN", "unit teste")

	guruLog := getGuruLog()

	guruLog.Warning(&log.LogWithFields{
		CustomerCode: "customerCode",
		Caller:       "TestLogWarn",
		InfoMessage: log.Fields{
			"msg 1": "mensagem 1",
			"msg 2": "mensagem 2",
			"msg 3": "mensagem 3",
		},
	}, "primeiro log",
	)

	fmt.Println("OK")
}

func TestLogPanic(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	log.InitLog("PANIC", "unit teste")

	guruLog := getGuruLog()

	guruLog.Panic(&log.LogWithFields{
		CustomerCode: "customerCode",
		Caller:       "TestLogPanic",
		InfoMessage: log.Fields{
			"msg 1": "mensagem 1",
			"msg 2": "mensagem 2",
			"msg 3": "mensagem 3",
		},
	}, "primeiro log",
	)

	fmt.Println("OK")
}
