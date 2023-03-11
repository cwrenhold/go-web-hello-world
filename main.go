package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	PageTitle string
	Value     string
}

func main() {
	http.HandleFunc("/", templateServer)
	log.Println("Starting on :8080")
	http.ListenAndServe(":8080", nil)
}

func templateServer(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Println(err)
		return
	}

	query := r.URL.Query()
	value := query.Get("value")

	pageData := PageData{
		PageTitle: "Hello World",
		Value:     value,
	}

	tmpl.Execute(w, pageData)
}
