package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func administerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("admin/administer.html")

	if err != nil {
		fmt.Fprintf(w, "Template error: %v", err)
		return
	}

	tmpl.Execute(w, nil)
}