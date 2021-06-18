package main

import (
	"gumSite_01/internal/app/servers"
	"log"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

func main() {
	config := servers.NewConfig()
	_, err := toml.DecodeFile("configs/apiserver.toml", config)
	if err != nil {
		log.Fatal(err)
	}
	if err := servers.Start(config); err != nil {
		log.Fatal(err)
	}
}
