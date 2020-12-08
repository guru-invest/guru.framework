package influx_connector

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type Influx struct {
	InfluxClient influxdb2.Client
}

// adicionar url e pass em variáveis
func InfluxClient() *Influx {
	client := influxdb2.NewClient("", "")

	connection := Influx{InfluxClient: client}

	return &connection
}

func SaveLog(measurement string, logData map[string]string) {
	client := InfluxClient()
	write := client.InfluxClient.WriteAPI("guru", "trade_audit_qa")
	newPoint := influxdb2.NewPointWithMeasurement(measurement)

	for k, v := range logData {
		newPoint.AddField(k, v)
	}
	newPoint.SetTime(time.Now())
	write.WritePoint(newPoint)

	CloseInfluxConnection(client)
}

func CloseInfluxConnection(connection *Influx) {
	connection.InfluxClient.Close()
}
