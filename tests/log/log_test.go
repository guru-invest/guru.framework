package log_test

import (
	"fmt"
	"testing"

	"github.com/guru-invest/guru.framework/src/helpers/log"
)

func TestLogInfo(t *testing.T) {
<<<<<<< HEAD
	log.InitLog("INFO")
=======
	log.InitLog()
>>>>>>> 455c96b115d6debae8dfe9c2bf8f98d50d150054

	guruLog := log.GuruLog{
		HTTPHeader: map[string][]string{},
	}
	guruLog.HTTPHeader.Add("user-agent", "Version | teste2 |")
	guruLog.HTTPHeader.Add("device-id", "ASD4AS56D4-4ASD54AS6D-4AS5D4AS65")
	guruLog.HTTPHeader.Add("client-ip", "192.168.0.1")
	guruLog.HTTPHeader.Add("service-name", "guru.framework")
	guruLog.HTTPHeader.Add("correlation-id", "AAAAA1-AAAAA2-AAAAA3")
	guruLog.HTTPHeader.Add("session-id", "FD45D84F7E")

<<<<<<< HEAD
	guruLog.Debug(log.LogWithFields{
=======
	guruLog.Error(log.LogWithFields{
>>>>>>> 455c96b115d6debae8dfe9c2bf8f98d50d150054
		CustomerCode: "customerCode",
		Message: log.Fields{
			"msg 1": "mensagem 1",
			"msg 2": "mensagem 2",
			"msg 3": "mensagem 3",
		},
	}, "primeiro log",
	)

	fmt.Println("OK")
}
