package main

type Product struct {
	ID          int
	Name        string
	Description string
	Cost        int
}

type Products struct {
	Data []Product
}

// temp vars used for reading the db
var id int
var name string
var desc string
var cost int

var lastKey int
