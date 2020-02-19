package controllers

import (
	"database/sql"
	"deadgopher/the_server/models"
	"encoding/json"
	"log"
	"net/http"
)


type user_controller struct {
	db *sql.DB
}

func NewUserController(db *sql.DB) *user_controller {
	uc := &user_controller{
		db,
	}

	return uc
}



func(uc *user_controller) Create(w http.ResponseWriter, r *http.Request){
	log.Print("In User Create Method!")

	var user_config models.UserConfig
	if err := json.NewDecoder(r.Body).Decode(&user_config); err != nil {
		response{
			Success: false,
			Payload: "Invalid Data",
			status_code:400,
		}.as_json(w)

		return
	}

	if valid, errs := user_config.Validate(); !valid {
		response{
			Success: false,
			Payload: errs,
			status_code:400,
		}.as_json(w)

		return
	}

	user := models.User(user_config, uc.db)
	if err := user.Save(); err != nil {
		response{
			Success: false,
			Payload: err.Error(),
			status_code:500,
		}.as_json(w)

		return
	}

	response{
		Success: true,
		Payload: nil,
		status_code:201,
	}.as_json(w)

}
func(uc *user_controller) Read(w http.ResponseWriter, r *http.Request){

}
func(uc *user_controller) Update(w http.ResponseWriter, r *http.Request){

}
func(uc *user_controller) Destroy(w http.ResponseWriter, r *http.Request){

}



