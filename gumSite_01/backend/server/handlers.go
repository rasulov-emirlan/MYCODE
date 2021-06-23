package server

import (
	"encoding/json"
	"fmt"
	"gumsite_01/storage"
	"net/http"
	"strconv"
)

// SelctProducts return a variable of type http.HandleFunct that
// Selects all products from database by calling storage.selectAllProducts
// it returns all products in json format
func (s *Server) SelectProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p, err := storage.SelectAllProducts(s.store)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, p)
	}
}

// SelctProductsByName return a variable of type http.HandleFunct that
// takes a (json string: name) and searches for a product with the same name in db
// by calling a storage.FindProductByName
// it returns that product in json format
func (s *Server) SelectProductByName() http.HandlerFunc {
	type Request struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		p, err := storage.FindProductByName(s.store, req.Name)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
		s.respond(w, r, http.StatusOK, p)
	}
}

func (s *Server) DeleteProduct() http.HandlerFunc {
	type request struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		storage.DeleteProduct(s.store, req.Name)
		s.respond(w, r, http.StatusOK, "done")
	}
}

func (s *Server) AddProduct() http.HandlerFunc {
	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Cost        string `json:"cost"`
		Image       string `json:"image"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		cost, err := strconv.Atoi(req.Cost)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
		p := storage.Product{
			Name:        req.Name,
			Description: req.Description,
			Cost:        cost,
			Image:       req.Image,
		}
		if err := storage.InserProduct(s.store, p); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
		s.respond(w, r, http.StatusCreated, p)
	}
}

func (s *Server) SelectAllOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		o, err := storage.SelectAllOrders(s.store)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		fmt.Fprint(w, o)
		s.respond(w, r, http.StatusOK, o)
	}
}

// InsertOrder returns a valiable of type http.HandleFunc
// that inserts a new order into the db table gumsite_orders
// by calling storage.InsertOrder
func (s *Server) InsertOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := storage.InsertOrder(s.store, storage.TestOrder()); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
		s.respond(w, r, http.StatusCreated, "done")
	}
}

// error is an analog of respond
// but used exucivly for sending errors back to the client
func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{
		"error": err.Error(),
	})
}

// respond sends a json data
// usualy used to send json back to the client
func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
