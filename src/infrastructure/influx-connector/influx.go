package influx_connector

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type Influx struct {
	Client influxdb2.Client
}

func (i Influx) InfluxConnection(url string, username string, password string, timeout uint) *Influx {
	client := influxdb2.NewClientWithOptions(url, username+":"+password, influxdb2.DefaultOptions().SetHTTPRequestTimeout(timeout))

	connection := Influx{Client: client}

	return &connection
}

func (i *Influx) SaveLog(database string, measurement string, customerCode string, ip string, eventType string, logData map[string]interface{}) {
	defer closeInfluxConnection(i)
	i.Client.
		WriteAPI("guru", database).
		WritePoint(
			influxdb2.NewPoint(
				measurement,
				map[string]string{"CustomerCode": customerCode, "IP": ip, "EventType": eventType},
				logData,
				time.Now(),
			),
		)
}

func closeInfluxConnection(connection *Influx) {
	connection.Client.Close()
}
