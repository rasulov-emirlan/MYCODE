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
var db *sql.DB

// all structs
type Product struct {
	ID   int
	Name string
	Cost int
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	var tID int
	var tName string
	var tCost int

	connStr := "user=postgres dbname=local_server host=cybertron sslmode=verify-full password=postgres"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected Db")

	rows, err := db.Query("SELECT * FROM gummy_products")
	if err != nil {
		log.Println("Error ocured while reading from db: ", err.Error())
	}

	var data []Product
	for rows.Next() {
		rows.Scan(&tID, &tName, &tCost)
		fmt.Println(tID, " ", tName, " ", tCost)
		data = append(data, Product{
			ID:   tID,
			Name: tName,
			Cost: tCost,
		})
	}
	tpl.ExecuteTemplate(w, "index.html", data)
}
