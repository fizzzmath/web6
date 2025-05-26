package main

import (
	"net/http"
	"net/http/cgi"
)

func main() {
	cgi.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		action := r.URL.Query().Get("action")

		if action == "administer" {
			administerHandler(w, r)
		} else {
			loginHandler(w, r)
		}
	}))
}