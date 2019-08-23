package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var Database *sql.DB

func Connect(DataBaseAddress string, DataBaseUsername string, DataBasePassword string, DataBaseName string){
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s", DataBaseUsername,
		DataBasePassword,  DataBaseName,  DataBaseAddress, "5432")
	db, err := sql.Open("postgres",dsn)

	if err != nil{
		fmt.Println("Error on connect to database ", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error on ping database ", err)
	}
	Database = db

}

func MapResult(rows *sql.Rows, queryName string) map[string][]map[string]interface{}{
	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Failed to map result", err)
	}
	m := make(map[string]interface{})
	mapped := make(map[string][]map[string]interface{})

	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			fmt.Println("Failed to map result", err)
		}
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		mapped[queryName] = append(mapped[queryName], m)
		m = make(map[string]interface{})
	}
	return mapped
}
