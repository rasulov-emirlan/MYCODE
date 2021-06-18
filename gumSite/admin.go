package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	// image uploading

	// 1. parse input, type multipart/form-data.
	r.ParseMultipartForm(10 << 20)

	// 2. retrieve file from posted form-data.
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error retrieving file from form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File size : %+v\n", handler.Size)
	fmt.Printf("MIME header: %+v\n", handler.Header)

	// 3. write temporary file on our server
	// tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	f, err := os.Create("static/images" + name + ".png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	img, _, _ := image.Decode(bytes.NewReader(fileBytes))
	err = png.Encode(f, img)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return
	}
	// ...............
	lastKey += 1
	_, err = db.Query("INSERT INTO gummy_products(id, name, description, cost) VALUES($1, $2, $3, $4);", lastKey, name, desc, cost)
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
	http.Redirect(w, r, "/", http.StatusSeeOther)
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
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
