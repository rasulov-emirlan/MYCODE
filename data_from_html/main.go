package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/save", save)
	http.ListenAndServe(":8080", nil)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

type Person struct {
	Name     string
	Password string
}

func save(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")

	person := &Person{Name: name, Password: password}

	b, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "save.html", person)
	fmt.Println(person)

	f, err := os.Open("save.json")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	f.Write(b)
	f.Close()
}
