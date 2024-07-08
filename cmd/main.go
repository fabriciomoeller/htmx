package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Login      string
			Senha      string
			urlDoc     string
			urlSuporte string
		}{
			Login:      "login",
			urlDoc:     "http://gurusistemas.com.br",
			urlSuporte: "http://gurubi.ddns.net:8070/",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Current time: " + time.Now().Format(time.RFC1123)))
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Login string
		}{
			Login: "login",
		}

		_ = tmpl.Execute(w, data)
	})

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
