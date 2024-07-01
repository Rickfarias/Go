package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
)

var templates *template.Template

type usuario struct {
	Nome string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {

	u := usuario {"Rick", "rickfarias@gmail.com",}

	templates.ExecuteTemplate(w, "home.html", u)
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home",  home)

	fmt.Println("Executando a porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}