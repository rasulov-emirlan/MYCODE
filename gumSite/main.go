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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/admin_page", adminPage)
	http.HandleFunc("/add_product", addProduct)
	http.HandleFunc("/delete_product", deleteProduct)

	http.ListenAndServe(":8080", nil)
	log.Println("Server is on!")
}

// all vars
var tpl *template.Template
var db *sql.DB
var connStr string

var errr error

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	connStr = "user=postgres password=postgres dbname=local_server sslmode=disable"
	db, errr = sql.Open("postgres", connStr)
	if errr != nil {
		log.Fatal(errr)
	}
	rows, err := db.Query("SELECT * FROM gummy_products")
	if err != nil {
		log.Panicln(err)
	}

	var Data Products
	for rows.Next() {
		rows.Scan(&id, &name, &desc, &cost)
		Data.Data = append(Data.Data, Product{
			ID:          id,
			Name:        name,
			Description: desc,
			Cost:        cost,
		})
	}
	lastKey = id
	defer db.Close()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	db, errr = sql.Open("postgres", connStr)
	if errr != nil {
		log.Fatal(errr)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM gummy_products")
	if err != nil {
		log.Panicln(err)
	}

	var Data Products
	for rows.Next() {
		rows.Scan(&id, &name, &desc, &cost)
		Data.Data = append(Data.Data, Product{
			ID:          id,
			Name:        name,
			Description: desc,
			Cost:        cost,
		})
	}
	lastKey = id
	fmt.Println(Data.Data)
	tpl.ExecuteTemplate(w, "index.html", Data)
}
