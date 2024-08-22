package database

import (
	"database/sql"
	"fmt"
)

func ConnectPostgres() {
	//connectionStr := "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	//
	//conn, err := sql.Open("postgres", connectionStr)
	//if err != nil {
	//	panic(err)
	//}
	//
	//conn.Close()
	connectionStr := "postgres://postgres:postgres@localhost:5432/sabervest?sslmode=disable"

	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	rows, err := conn.Query("SELECT matter();")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var version string
		rows.Scan(&version)
		fmt.Println(version)
	}

	rows.Close()

	conn.Close()
}

// This time the global variable is unexported.
//var db *sql.DB

//func InitDB(dataSourceName string) error {
//	var err error
//
//	db, err = sql.Open("postgres", dataSourceName)
//	if err != nil {
//		return err
//	}
//
//	return db.Ping()
//}
