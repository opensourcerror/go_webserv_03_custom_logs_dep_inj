package main

import (
	"html/template"
	"net/http"
)

// http://localhost:4000/sb
func (app *application) sb(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/pages/secondBreakfast.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "secondBreakfast page not found", http.StatusInternalServerError)
		return
	}

	// "sb" IS NOT THE FILE NAME
	// it's the name you defined inside the template? {{define "sb"}}
	err = ts.ExecuteTemplate(w, "sb", nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "'sb' template not found", http.StatusInternalServerError)
		return
	}

}
