package main

import (
	"encoding/json"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/go-chi/chi/v5"
	"github.com/wederer/go-chi-demo/internal/models"
	"log"
	"net/http"
)

func (s *Server) HelloWorld(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello World!"))
}

func (s *Server) GetBook(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	book, err := s.Books.Read(idParam)
	if book == nil || driver.IsNotFoundGeneral(err) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		log.Printf("Failed to read document: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Marshall(book, w)
}

func (s *Server) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	err := s.Books.Delete(idParam)
	if err != nil {
		if err.Error() == models.NotFoundErrorCode {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func (s *Server) GetBooks(w http.ResponseWriter, _ *http.Request) {
	books, err := s.Books.ReadAll()

	if err != nil {
		if err.Error() == models.NotFoundErrorCode {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Marshall(books, w)
}

func (s *Server) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Printf("Failed to parse body %v with error: %v", r.Body, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := s.Books.Create(book)
	if err != nil {
		log.Printf("Failed to create document: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Marshall(result, w)
}

func (s *Server) GetProtectedData(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(fmt.Sprint("Secret Protected Info")))
}

func (s *Server) GetAdminInfo(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(fmt.Sprint("Even more secret info that only admins can read.")))
}
