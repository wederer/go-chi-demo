package models

import "github.com/arangodb/go-driver"

type Book struct {
	Key     string `json:"id"`
	Title   string `json:"title"`
	NoPages int    `json:"no_pages"`
}

//type ModelInterface[T Book | any] interface {
//	Read(key string) (T, error)
//	ReadAll() ([]T, error)
//	Delete(key string) error
//	Create(T) (T, error)
//}
//type Model[T any] struct {
//	Read func(key string) (T, error)
//	ReadAll() ([]T, error)
//	Delete(key string) error
//	Create(T) (T, error)
//}
//func (m Model[T]) Read(key string) (T, error) {
//	return m.Read(key)
//}

type BookModelInterface interface {
	Read(key string) (*Book, error)
	ReadAll() ([]Book, error)
	Delete(key string) error
	Create(Book) (Book, error)
}

type BookModel struct {
	db driver.Database
}

func (b *BookModel) ReadAll() ([]Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BookModel) Delete(key string) error {
	//TODO implement me
	panic("implement me")
}

func (b *BookModel) Create(book Book) (Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BookModel) Read(key string) (*Book, error) {
	booksCollection, err := b.db.Collection(nil, "books")
	if err != nil {
		return nil, err
	}

	var book Book
	meta, err := booksCollection.ReadDocument(nil, key, &book)
	if err != nil {
		return nil, err
	}

	if &book != nil {
		book.Key = meta.Key
		return &book, nil
	}
	panic("some unknown error")
}

func New(dbIn driver.Database) *BookModel {
	b := &BookModel{
		db: dbIn,
	}
	return b
}
