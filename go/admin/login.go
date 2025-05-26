package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("admin/login.html")

	if err != nil {
		fmt.Fprintf(w, "Template error: %v", err)
	}

	tmpl.Execute(w, nil)
}