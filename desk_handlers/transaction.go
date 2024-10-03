package handlers

import (
	"desk_objects/objects"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request, data objects.Login) {
	template_parser := template.Must(template.ParseFiles("static/templates/home.html", "static/templates/logins.html"))
	if data != nil {
		template_parser.ExecuteTemplate(w, "home.html", data)
		return
	}
	template_parser.ExecuteTemplate(w, "home.html", nil)
	return
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

}
