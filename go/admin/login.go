package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type (
	Admin struct {
		Login string
		Password string
	}
)

func valid(admin Admin) (bool, error) {
	db, err := sql.Open("mysql", "u68867:6788851@/u68867")

	if err != nil {
		return false, err
	}

	defer db.Close()

	login, password := "", ""

	sel, err := db.Query(`
		SELECT Login, Password 
		FROM ADMIN
	`)

	if err != nil {
		return false, err
	}

	defer sel.Close();

	for sel.Next() {
		err := sel.Scan(&login, &password)

		if err != nil {
			return false, err
		}
	}

	if login != admin.Login || fmt.Sprintf("%x", password) != admin.Password {
		return false, nil
	}

	return true, nil
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

		ok, err := valid(admin)

		if err != nil {
			fmt.Fprintf(w, "MySQL error: %v", err)
			return
		}

		if !ok {
			tmpl.Execute(w, "Неверные логин или пароль")
			return
		}
	}

	tmpl.Execute(w, nil)
}