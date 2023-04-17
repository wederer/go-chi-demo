package models

import (
	"errors"
	"github.com/arangodb/go-driver"
	"log"
)

type Book struct {
	Key     string `json:"id"`
	Title   string `json:"title"`
	NoPages int    `json:"no_pages"`
}

type BookModel struct {
	db   driver.Database
	coll driver.Collection
}

const NotFoundErrorCode = "not found"

func (m *BookModel) ReadAll() ([]*Book, error) {
	var books = make(map[string]interface{}, 0)
	var returnSlice = make([]*Book, 0)
	ctx := driver.WithQueryCount(nil)
	cursor, err := m.db.Query(ctx, "FOR doc IN books RETURN doc", books)
	defer cursor.Close()
	if err != nil {
		return nil, err
	}

	for i := 0; i < int(cursor.Count()); i++ {
		var book *Book
		meta, _ := cursor.ReadDocument(nil, &book)
		book.Key = meta.Key
		returnSlice = append(returnSlice, book)
	}

	if &books == nil || driver.IsNotFoundGeneral(err) {
		return nil, errors.New(NotFoundErrorCode)
	}

	if err != nil {
		log.Printf("Failed to read document: %v", err)
		return nil, err
	}

	return returnSlice, nil
}

func (m *BookModel) Delete(key string) error {
	_, err := m.coll.RemoveDocument(nil, key)
	if driver.IsNotFoundGeneral(err) {
		return errors.New(NotFoundErrorCode)
	}
	if err != nil {
		log.Printf("Failed to read document: %v", err)
		return err
	}

	return nil
}

func (m *BookModel) Create(book Book) (*Book, error) {
	var result Book
	ctx := driver.WithReturnNew(nil, &result)
	_, err := m.coll.CreateDocument(ctx, &book)
	if err != nil {
		log.Printf("Failed to create document: %v", err)
		return nil, err
	}

	log.Printf("Created Book '%s'\n", book.Title)
	return &result, nil
}

func (m *BookModel) Read(key string) (*Book, error) {
	var book *Book
	meta, err := m.coll.ReadDocument(nil, key, &book)
	if err != nil {
		return nil, err
	}

	if book != nil {
		book.Key = meta.Key
		return book, nil
	}
	panic("some unknown error")
}

func NewBookModel(dbIn driver.Database) *BookModel {
	collectionName := "books"
	colExists, err := dbIn.CollectionExists(nil, collectionName)
	if err != nil {
		log.Fatalf("Failed to check for collection: %v", err)
	}
	var coll driver.Collection
	if colExists == true {
		coll, err = dbIn.Collection(nil, collectionName)
	} else {
		coll, err = dbIn.CreateCollection(nil, collectionName, nil)
	}
	if err != nil {
		log.Fatalf("Failed to create/get collection: %v", err)
	}

	return &BookModel{
		db:   dbIn,
		coll: coll,
	}
}
