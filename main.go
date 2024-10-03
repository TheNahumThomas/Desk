package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	"static\\templates\\base.html",
	"static\\templates\\dash.html",
))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		data := map[string]interface{}{
			"Title": "Dashboard",
			"Name":  "Nahum",
		}

		err := tmpl.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Start the server on port 8080
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
