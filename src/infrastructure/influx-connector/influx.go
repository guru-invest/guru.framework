package influx_connector

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type Influx struct {
	InfluxClient influxdb2.Client
}

func InfluxClient(url string, username string, password string) *Influx {
	client := influxdb2.NewClient(url, username+":"+password)

	connection := Influx{InfluxClient: client}

	return &connection
}

func (client *Influx) SaveLog(database string, measurement string, logData map[string]string) {
	defer closeInfluxConnection(client)
	write := client.InfluxClient.WriteAPI("guru", database)
	newPoint := influxdb2.NewPointWithMeasurement(measurement)

	for k, v := range logData {
		newPoint.AddField(k, v)
	}
	newPoint.SetTime(time.Now())
	write.WritePoint(newPoint)

}

func closeInfluxConnection(connection *Influx) {
	connection.InfluxClient.Close()
}
