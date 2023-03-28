package main

import (
	"encoding/json"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Book struct {
	Title   string `json:"title"`
	NoPages int    `json:"no_pages"`
}

func (s *Server) HelloWorld(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello World!"))
}

func (s *Server) GetBook(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	var book Book
	_, err := s.Books.ReadDocument(nil, idParam, &book)
	if &book == nil || driver.IsNotFoundGeneral(err) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		log.Printf("Failed to read document: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%v", book)))
}

func (s *Server) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Printf("Failed to parse body %v with error: %v", r.Body, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var result Book
	ctx := driver.WithReturnNew(nil, &result)
	_, err = s.Books.CreateDocument(ctx, book)
	if err != nil {
		log.Printf("Failed to create document: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Created Book '%s'\n", s.Books.Name())
	w.Write([]byte(fmt.Sprintf("%v", result)))
}

func (s *Server) GetProtectedData(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(fmt.Sprint("Secret Protected Info")))
}

func (s *Server) GetAdminInfo(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(fmt.Sprint("Even more secret info that only admins can read.")))
}
