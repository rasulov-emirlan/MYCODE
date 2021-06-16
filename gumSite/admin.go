package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
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

	rows, err := db.Query("")
	if err != nil {
		log.Println("Error occured while reading db:", err)
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
	fmt.Println(Data.Data)
	tpl.ExecuteTemplate(w, "index.html", Data)
}
