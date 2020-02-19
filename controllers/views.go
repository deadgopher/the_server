package controllers

import (
	"net/http"
	"html/template"
)

type view_controller struct {
	views *template.Template
}

func NewViewController() *view_controller {
	tmp := template.Must(template.ParseGlob("views/*.gohtml"))
	vc := &view_controller{views:tmp}

	return vc
}

func(vc *view_controller) Index(w http.ResponseWriter, r *http.Request){
	vc.views.ExecuteTemplate(w, "index", nil)
}
