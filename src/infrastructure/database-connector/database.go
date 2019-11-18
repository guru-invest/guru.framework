package database_connector

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

type DataBase struct{
	URI string
	Connection *gorm.DB
}

func (c *DataBase) ConnectToDatabase(url string, username string, password string, dbname string, host string, port string) (*DataBase, error){
	dsn := fmt.Sprintf(url, username,
		password, dbname, host, port)
	db, err := gorm.Open("postgres", dsn)
	database := DataBase{
		URI: dsn,
		Connection:db,
	}
	if err != nil {
		return nil, errors.Wrap(err, "error on connecting to database")
	}
	return &database, nil
}

func (c *DataBase) DisconnectFromDatabase(database *DataBase) error {
	err := database.Connection.Close()
	if err != nil {
		return errors.Wrap(err, "error on disconnecting from database")
	}
	return nil
}
