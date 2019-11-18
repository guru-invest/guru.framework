package infrastructure

import (
	database_connector "github.com/guru-invest/guru.framework/src/infrastructure/database-connector"
	http_connector "github.com/guru-invest/guru.framework/src/infrastructure/http-connector"
)

type ConnectorService struct{
	HttpClient *http_connector.HttpClient
	Database *database_connector.DataBase
}

func (cs *ConnectorService) NewDatabaseConnectorService(url string, username string, password string, dbname string, host string, port string) *database_connector.DataBase{
	cs.Database, _ = cs.Database.ConnectToDatabase(url, username, password, dbname, host, port)
	return cs.Database
}

func (cs *ConnectorService) NewHttpConnectorService() *http_connector.HttpClient{
	return cs.HttpClient
}
