package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
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
	for rows.Next() {
		rows.Scan(&id, &name, &cost)
		fmt.Println(id, " ", name, " ", cost)
	}

}
