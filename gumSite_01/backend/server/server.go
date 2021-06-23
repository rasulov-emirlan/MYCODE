package server

import (
	"database/sql"
	"gumsite_01/storage"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

type Server struct {
	r     *mux.Router
	store storage.Storage
}

func NewServer() *Server {
	s := &Server{
		r: mux.NewRouter(),
	}
	return s
}

func (s *Server) Start(port string, dbConfig string) error {
	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.store = storage.New(db)
	defer db.Close()
	s.handleRequests()
	return http.ListenAndServe(port, s.r)
}

func (s *Server) handleRequests() {
	s.r.HandleFunc("/selectAllCategories", s.SelectAllCategories()).Methods("GET")
	s.r.HandleFunc("/selectCategoryByName", s.SelectCategoryByName()).Methods("GET")
	s.r.HandleFunc("/insertCategory", s.InsertCategory()).Methods("POST")
	s.r.HandleFunc("/deleteCategory", s.DeleteCategory()).Methods("DELETE")

	// handles all functions for product assortiment management
	s.r.HandleFunc("/selectAllProducts", s.SelectProducts()).Methods("GET")
	s.r.HandleFunc("/selectProductByName", s.SelectProductByName()).Methods("GET")
	s.r.HandleFunc("/insertProduct", s.AddProduct()).Methods("POST")
	s.r.HandleFunc("/deleteProduct", s.DeleteProduct()).Methods("DELETE")

	// handles all functions for order management
	s.r.HandleFunc("/selectAllOrders", s.SelectAllOrders()).Methods("GET")
	s.r.HandleFunc("/selectOrderByName", s.SelectOrderByName()).Methods("GET")
	s.r.HandleFunc("/insertOrder", s.InsertOrder()).Methods("POST")
	s.r.HandleFunc("/deleteOrder", s.DeleteOrder()).Methods("DELETE")

}
