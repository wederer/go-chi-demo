package main

import (
	"fmt"
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

	if _, err := s.Books.ReadDocument(nil, idParam, &book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to read document: %v", err)
	}
	if &book == nil {
		fmt.Println("Book not found")
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write([]byte(fmt.Sprintf("%v", book)))
}

func (s *Server) CreateBook(w http.ResponseWriter, r *http.Request) {
	book := Book{
		Title:   "ArangoDB Cookbook",
		NoPages: 257,
	}
	meta, err := s.Books.CreateDocument(nil, book)
	if err != nil {
		log.Printf("Failed to create document: %v", err)
	}
	fmt.Printf("Created document in collection '%s'\n", s.Books.Name())

	var result Book
	if _, err := s.Books.ReadDocument(nil, meta.Key, &result); err != nil {
		log.Printf("Failed to read document: %v", err)
	}
	fmt.Printf("Read book '%+v'\n", result)
}
