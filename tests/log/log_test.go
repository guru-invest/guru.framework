package log_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/guru-invest/guru.framework/src/helpers/log"
)

func TestLogInfo(t *testing.T) {
	log.InitLog("INFO", "unit teste")

	guruLog := log.GuruLog{
		HTTPHeader: &http.Header{},
	}
	guruLog.HTTPHeader.Add("user-agent", "Version | teste2 |")
	guruLog.HTTPHeader.Add("device-id", "ASD4AS56D4-4ASD54AS6D-4AS5D4AS65")
	guruLog.HTTPHeader.Add("X-Forwarded-For", "192.168.0.1")
	guruLog.HTTPHeader.Add("service-name", "guru.framework")
	guruLog.HTTPHeader.Add("correlation-id", "AAAAA1-AAAAA2-AAAAA3")
	guruLog.HTTPHeader.Add("session-id", "FD45D84F7E")

	guruLog.Error(&log.LogWithFields{
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

func TestLogInfoFieldsNull(t *testing.T) {
	log.InitLog("INFO", "unit teste")

	guruLog := log.GuruLog{
		HTTPHeader: nil,
	}

	guruLog.Info(nil, "primeiro log")

	fmt.Println("OK")
}
