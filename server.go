package main

import (
	"log"
	"net/http"
	"text/template"
)

var WebHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	templateName := r.FormValue("template")
	templateData, ok := CurrentConfig.Templates[templateName]
	if !ok {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	message, help := r.FormValue("message"), r.FormValue("help")
	if message != "" {
		templateData.Message = message
	}

	if help != "" {
		templateData.Help = help
	}

	template, err := template.ParseFiles("index.gohtml")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = template.Execute(w, templateData)
	if err != nil {
		log.Println(err)
		return
	}
})
