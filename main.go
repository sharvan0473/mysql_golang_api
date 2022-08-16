package main

import (
	"sharvan/mux/api/app"
	"sharvan/mux/api/data_source/db"
)

func main() {
	db.ConnectDB()
	app.StartApplication()

}
