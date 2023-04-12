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

//type Model[T any] struct {
//	Read func(key string) (T, error)
//	ReadAll() ([]T, error)
//	Delete(key string) error
//	Create(T) (T, error)
//}
//func (m Model[T]) Read(key string) (T, error) {
//	return m.Read(key)
//}

type BookModel struct {
	db   driver.Database
	coll driver.Collection
}

const NotFoundErrorCode = "not found"

func (b *BookModel) ReadAll() ([]Book, error) {
	var books = make(map[string]interface{}, 0)
	var booksSlice = make([]Book, 0)
	ctx := driver.WithQueryCount(nil)
	cursor, err := b.db.Query(ctx, "FOR doc IN books RETURN doc", books)
	defer cursor.Close()
	if err != nil {
		return nil, err
	}

	for i := 0; i < int(cursor.Count()); i++ {
		var book Book
		meta, _ := cursor.ReadDocument(nil, &book)
		book.Key = meta.Key
		booksSlice = append(booksSlice, book)
	}

	if &books == nil || driver.IsNotFoundGeneral(err) {
		return nil, errors.New(NotFoundErrorCode)
	}

	if err != nil {
		log.Printf("Failed to read document: %v", err)
		return nil, err
	}

	return booksSlice, nil
}

func (b *BookModel) Delete(key string) error {
	_, err := b.coll.RemoveDocument(nil, key)
	if driver.IsNotFoundGeneral(err) {
		return errors.New(NotFoundErrorCode)
	}
	if err != nil {
		log.Printf("Failed to read document: %v", err)
		return err
	}

	return nil
}

func (b *BookModel) Create(book Book) (*Book, error) {
	var result Book
	ctx := driver.WithReturnNew(nil, &result)
	_, err := b.coll.CreateDocument(ctx, book)
	if err != nil {
		log.Printf("Failed to create document: %v", err)
		return nil, err
	}

	log.Printf("Created Book '%s'\n", book.Title)
	return &result, nil
}

func (b *BookModel) Read(key string) (*Book, error) {
	var book Book
	meta, err := b.coll.ReadDocument(nil, key, &book)
	if err != nil {
		return nil, err
	}

	if &book != nil {
		book.Key = meta.Key
		return &book, nil
	}
	panic("some unknown error")
}
