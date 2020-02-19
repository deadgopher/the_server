package util

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1)/the_server")
	if err != nil {
		log.Fatal("!!! ERROR CONNECTING TO DB !!!", err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Ping Error", err.Error())
	}

	return db
}
