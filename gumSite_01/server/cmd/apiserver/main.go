package main

import (
	"gumSite_01/internal/app/servers"
	"log"
)

func main() {
	if err := servers.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
