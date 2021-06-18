package servers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gumSite_01/internal/app/database/products"
	"net/http"

	"github.com/gorilla/mux"
)

// server ...
type server struct {
	router *mux.Router
	db     *sql.DB
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer() server {
	s := server{
		router: mux.NewRouter(),
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/", s.homePage()).Methods("GET")
	s.router.HandleFunc("/getAll", s.getAll()).Methods("GET")
	s.router.HandleFunc("/findByName", s.homePage()).Methods("GET")
	s.router.HandleFunc("/insertProduct", s.homePage()).Methods("POST")
}

func (s *server) homePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, "Home Page of REST_01")
	}
}

func (s *server) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p, err := products.GetALL(s.db)
		if err != nil {
			fmt.Fprint(w, "error searching db")
		}

		s.respond(w, r, http.StatusOK, p)
	}
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
