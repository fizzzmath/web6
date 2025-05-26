package main

import (
	"crypto/sha256"
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

func valid(admin Admin, w http.ResponseWriter) (bool, error) {
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

	if login != admin.Login || fmt.Sprintf("%x", sha256.Sum256([]byte(password))) != admin.Password {
		fmt.Fprintf(w, "Ожидалось:\t%s\t%s\nПолучено:\t%s\t%x", login, sha256.Sum256([]byte(password)), admin.Login, admin.Password)
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

		ok, err := valid(admin, w)
		return

		if err != nil {
			fmt.Fprintf(w, "MySQL error: %v", err)
			return
		}

		if !ok {
			tmpl.Execute(w, "Неверные логин или пароль")
			return
		}

		tmpl, err = template.ParseFiles("admin/administrate.html")

		if err != nil {
			fmt.Fprintf(w, "Template error: %v", err)
			return
		}
	}

	tmpl.Execute(w, nil)
}