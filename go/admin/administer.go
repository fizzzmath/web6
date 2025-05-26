package main

import (
	"fmt"
	"net/http"
)

func administerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Добро пожаловать на страницу входа!")
}