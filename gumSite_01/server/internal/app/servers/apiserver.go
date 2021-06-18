package servers

import (
	"database/sql"
	"log"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	s := NewServer()
	return http.ListenAndServe(config.BindAddr, s.router)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
