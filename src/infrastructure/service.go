package infrastructure

import (
	database_connector "github.com/guru-invest/guru.framework/src/infrastructure/database-connector"
	http_connector "github.com/guru-invest/guru.framework/src/infrastructure/http-connector"
	influx_connector "github.com/guru-invest/guru.framework/src/infrastructure/influx-connector"
)

type ConnectorService struct {
	HttpClient *http_connector.HttpClient
	Database   *database_connector.DatabaseConnector
	Influx     *influx_connector.Influx
}

// func (cs *ConnectorService) NewDatabaseConnectorService(url string, username string, password string, dbname string, host string, port string) *database_connector.DatabaseConnection {
// 	cs.Database, _ = cs.Database.ConnectForServless(url, username, password, dbname, host, port)
// 	return cs.Database
// }

func (cs *ConnectorService) NewInfluxConnectorService(url string, username string, password string, timeout uint) *influx_connector.Influx {
	cs.Influx = cs.Influx.InfluxConnection(url, username, password, timeout)
	return cs.Influx
}

func (cs *ConnectorService) NewHttpConnectorService() *http_connector.HttpClient {
	return cs.HttpClient
}
