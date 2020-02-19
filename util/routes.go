package util

import (
	"database/sql"
	"deadgopher/the_server/controllers"
	"github.com/gorilla/mux"
)



func Router(db *sql.DB) *mux.Router {

	views := controllers.NewViewController()
	users := controllers.NewUserController(db)


	r := mux.NewRouter();

	r.HandleFunc("/", views.Index)

	r.Methods("post").Path("/user").HandlerFunc(users.Create)


	return r
}