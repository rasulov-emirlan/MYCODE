package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func adminPage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "admin_page.html", nil)
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	db, errr = sql.Open("postgres", connStr)
	if errr != nil {
		log.Fatal(errr)
	}
	defer db.Close()
	name = r.FormValue("name")
	desc = r.FormValue("description")
	cost, errr = strconv.Atoi(r.FormValue("cost"))
	if errr != nil {
		log.Println("Error ocured while getting values from forms: ", errr)
	}
	lastKey += 2
	_, err := db.Query("INSERT INTO gummy_products(id, name, description, cost) VALUES($1, $2, $3, $4);", lastKey, name, desc, cost)
	if err != nil {
		log.Println("Error occured while reading db:", err)
	}
	rows, err := db.Query("SELECT name, description, cost FROM gummy_products")
	if err != nil {
		log.Println("Error occured while reading db:", err)
	}
	var Data Products
	for rows.Next() {
		rows.Scan(&name, &desc, &cost)
		Data.Data = append(Data.Data, Product{
			Name:        name,
			Description: desc,
			Cost:        cost,
		})
	}
	lastKey = id
	fmt.Println(Data.Data)
	tpl.ExecuteTemplate(w, "index.html", Data)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	db, errr = sql.Open("postgres", connStr)
	if errr != nil {
		log.Fatal(errr)
	}
	defer db.Close()
	name = r.FormValue("name")
	cost, errr = strconv.Atoi(r.FormValue("cost"))
	if errr != nil {
		log.Println("Error ocured while getting values from forms: ", errr)
	}
	lastKey += 2
	_, err := db.Query("DELETE FROM gummy_products WHERE name=$1;", name)
	if err != nil {
		log.Println("Error occured while reading db:", err)
	}
	rows, err := db.Query("SELECT name, description, cost FROM gummy_products")
	if err != nil {
		log.Println("Error occured while reading db:", err)
	}
	var Data Products
	for rows.Next() {
		rows.Scan(&name, &desc, &cost)
		Data.Data = append(Data.Data, Product{
			Name:        name,
			Description: desc,
			Cost:        cost,
		})
	}
	lastKey = id
	fmt.Println(Data.Data)
	tpl.ExecuteTemplate(w, "index.html", Data)
}
