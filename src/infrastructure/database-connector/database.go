package database_connector

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase struct {
	URI        string
	Connection *gorm.DB
}

func (c *DataBase) ConnectToDatabase(url string, username string, password string, dbname string, host string, port string) (*DataBase, error) {
	dsn := fmt.Sprintf(url, username,
		password, dbname, host, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	database := DataBase{
		URI:        dsn,
		Connection: db,
	}
	if err != nil {
		return nil, errors.Wrap(err, "error on connecting to database")
	}
	return &database, nil
}
