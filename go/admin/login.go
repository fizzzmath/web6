package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type (
	Admin struct {
		Login string
		Password string
	}
)

func valid(admin Admin) bool {
	return false
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("admin/login.html")

	if err != nil {
		fmt.Fprintf(w, "Template error: %v", err)
	}

	if r.Method == http.MethodPost {
		admin := Admin {
			Login: r.FormValue("login"),
			Password: r.FormValue("password"),
		}

		if !valid(admin) {
			tmpl.Execute(w, "Неверные логин или пароль")
			return
		}
	}

	tmpl.Execute(w, nil)
}