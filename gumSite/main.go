package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
	log.Println("Server is on!")
}

// all vars
var tpl *template.Template

// all structs
type Product struct {
	ID   int
	Name string
	Cost int
}

type Products struct {
	Data []Product
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	connStr := "user=postgres password=postgres dbname=local_server sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM gummy_products")
	if err != nil {
		log.Panicln(err)
	}

	var id int
	var name string
	var cost int

	var Data Products
	for rows.Next() {
		rows.Scan(&id, &name, &cost)
		Data.Data = append(Data.Data, Product{
			ID:   id,
			Name: name,
			Cost: cost,
		})
	}
	fmt.Println(Data.Data)
	tpl.ExecuteTemplate(w, "index.html", Data)
}
