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

	book, err := s.Books2.Read(idParam)
	if book == nil || driver.IsNotFoundGeneral(err) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		log.Printf("Failed to read document: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%v", *book)))
}

func (s *Server) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	_, err := s.Books.RemoveDocument(nil, idParam)
	if driver.IsNotFoundGeneral(err) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		log.Printf("Failed to read document: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func (s *Server) GetBooks(w http.ResponseWriter, r *http.Request) {
	var books = make(map[string]interface{}, 0)
	var booksSlice = make([]models.Book, 0)
	ctx := driver.WithQueryCount(nil)
	cursor, err := s.DB.Query(ctx, "FOR doc IN books RETURN doc", books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close()
	for i := 0; i < int(cursor.Count()); i++ {
		var book models.Book
		meta, _ := cursor.ReadDocument(nil, &book)
		book.Key = meta.Key
		booksSlice = append(booksSlice, book)
	}

	if &books == nil || driver.IsNotFoundGeneral(err) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		log.Printf("Failed to read document: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mJson, err := json.Marshal(booksSlice)

	if err != nil {
		log.Printf("Failed to marshall document: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(mJson)
}

func (s *Server) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Printf("Failed to parse body %v with error: %v", r.Body, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var result models.Book
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
