package main

import (
	"deadgopher/the_server/util"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main(){

	port := ":8000"

	db := util.Connect()
	defer db.Close()

	static := http.StripPrefix("/static/", http.FileServer(http.Dir("./views/static")))
	router := util.Router(db)

	http.Handle("/static/", static)
	http.Handle("/", router)

	log.Printf("Listining on port %v...", port)
	http.ListenAndServe(port, nil)
}