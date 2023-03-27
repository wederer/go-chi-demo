package main

import (
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

func (s *Server) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func (s *Server) GetBook(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	var book Book
	_, err := s.Books.ReadDocument(nil, idParam, &book)
	if &book == nil || driver.IsNotFoundGeneral(err) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Book with id %v not found.", idParam)))
		return
	}
	if err != nil {
		log.Printf("Failed to read document: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Write([]byte(fmt.Sprintf("%v", book)))
}

func (s *Server) CreateBook(w http.ResponseWriter, r *http.Request) {
	book := Book{
		Title:   "ArangoDB Cookbook",
		NoPages: 257,
	}
	var result Book
	ctx := driver.WithReturnNew(nil, &result)
	_, err := s.Books.CreateDocument(ctx, book)
	if err != nil {
		log.Printf("Failed to create document: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	fmt.Printf("Created document in collection '%s'\n", s.Books.Name())

	w.Write([]byte(fmt.Sprintf("%v", result)))
}
