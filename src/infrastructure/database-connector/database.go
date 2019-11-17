package database_connector

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

func ConnectToDatabase(url string, username string, password string, dbname string, host string, port string) (*gorm.DB, error){
	dsn := fmt.Sprintf(url, username,
		password, dbname, host, port)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "error on connecting to database")
	}
	return db, nil
}

func DisconnectFromDatabase(database *gorm.DB) error {
	err := database.Close()
	if err != nil {
		return errors.Wrap(err, "error on disconnecting from database")
	}
	return nil
}
