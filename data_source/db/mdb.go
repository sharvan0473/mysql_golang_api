package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	ProjectDB *sql.DB
)

func ConnectDB() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "root", "127.0.0.1", "sample_project")

	var err error
	ProjectDB, err = sql.Open("mysql", dataSourceName)
	if err != nil {

		panic(err)
	}
	fmt.Println("Database Connected successfully")
}
