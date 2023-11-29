package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "avnadmin:AVNS_6Qd8Xrn4YPxSC9IFNmc@tcp(projetotcc-gabrielcaetanounivicosa-13b9.a.aivencloud.com:23847)/projeto_tcc")
	if err != nil {
		panic(err)
	}

	log.Println("Database connected")
	DB = db
}
