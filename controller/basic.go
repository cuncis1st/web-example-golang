package controller

import (
	"html/template"
	"net/http"
)

func Basic() {
	var tmpl, err = template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
		return
	}

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]interface{})
		data["name"] = "Batman"
		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]interface{})
		data["name"] = "Superman"
		data["key"] = "pR9TXG35w1vscG7FOddLWKTNtTaN7d1bIgLd"

		err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
