package database_connector

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type DatabaseConnector struct {
	DataSource struct {
		Port     int
		URL      string
		Username string
		Password string
		Database string
	}
	ConfigPool struct {
		SetConnMaxLifetime time.Duration
	}
}

func (t DatabaseConnector) ConnectForServless(databaseOption DatabaseConnector) (*sql.DB, error) {

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d timezone=%s",
		databaseOption.DataSource.Username,
		databaseOption.DataSource.Password,
		databaseOption.DataSource.Database,
		databaseOption.DataSource.URL,
		databaseOption.DataSource.Port,
		"America/Sao_Paulo")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "error on connecting to database")
	}
	db.SetMaxOpenConns(1)                                               // maximo de nova conexao por pool de conexao
	db.SetMaxIdleConns(0)                                               //maximo de conexão inativa aguardando reuso
	db.SetConnMaxLifetime(databaseOption.ConfigPool.SetConnMaxLifetime) //tempo maximo para expirar uma conexao

	return db, nil
}

// type DataBase struct {
// 	URI        string
// 	Connection *gorm.DB
// }

// func (c *DataBase) ConnectToDatabase(url string, username string, password string, dbname string, host string, port string) (*DataBase, error) {
// 	dsn := fmt.Sprintf(url, username,
// 		password, dbname, host, port)
// 	db, err := gorm.Open("postgres", dsn)
// 	database := DataBase{
// 		URI:        dsn,
// 		Connection: db,
// 	}
// 	if err != nil {
// 		return nil, errors.Wrap(err, "error on connecting to database")
// 	}
// 	return &database, nil
// }

// func (c *DataBase) DisconnectFromDatabase(database *DataBase) error {
// 	err := database.Connection.Close()
// 	if err != nil {
// 		return errors.Wrap(err, "error on disconnecting from database")
// 	}
// 	return nil
// }
