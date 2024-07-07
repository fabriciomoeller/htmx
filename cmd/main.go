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
			Message string
		}{
			Message: "Hello World",
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

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
