package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Success bool `json:"success"`
	Payload interface{} `json:"payload"`
	status_code int
}

func (r response) as_json(w http.ResponseWriter) {
	log.Printf("Sending response: (%v : %v)", r.Success, r.Payload)
	if r.status_code != 0 {w.WriteHeader(r.status_code)}
	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(r)
	if err != nil {
		panic(err.Error())
	}
	w.Write(js)
}