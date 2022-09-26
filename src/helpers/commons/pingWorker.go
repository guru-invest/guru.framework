package commons

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type pingWorker struct {
	_endpoint string
	_isLoaded bool
}

var worker pingWorker
var start time.Time

func (t *pingWorker) makeHttpRequest() bool {
	req, _ := http.NewRequest("POST", t._endpoint, nil)
	req.Header.Add("tenant", "guru")

	start = time.Now()
	resp, err := http.DefaultClient.Do(req)

	log.Debug(fmt.Sprintf("Ping response took %v", time.Since(start)))
	return err == nil && resp.StatusCode == 200
}

func ping(url string) bool {
	if worker._isLoaded {
		worker._endpoint = url
	}
	return worker.makeHttpRequest()
}

func StartPingWorker(url string) {
	worker = pingWorker{
		_isLoaded: true,
	}

	for {
		ping(url)
		time.Sleep(time.Duration(10) * time.Second)
	}
}
