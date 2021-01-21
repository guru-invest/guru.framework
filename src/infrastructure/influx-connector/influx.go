package influx_connector

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type Influx struct {
	Client influxdb2.Client
}

func (i Influx) InfluxConnection(url string, username string, password string) *Influx {
	client := influxdb2.NewClient(url, username+":"+password)

	connection := Influx{Client: client}

	return &connection
}

func (i *Influx) SaveLog(database string, measurement string, logData map[string]string) {
	defer closeInfluxConnection(i)
	write := i.Client.WriteAPI("guru", database)
	newPoint := influxdb2.NewPointWithMeasurement(measurement)

	for k, v := range logData {
		newPoint.AddField(k, v)
	}
	newPoint.SetTime(time.Now())
	write.WritePoint(newPoint)

}

func closeInfluxConnection(connection *Influx) {
	connection.Client.Close()
}
