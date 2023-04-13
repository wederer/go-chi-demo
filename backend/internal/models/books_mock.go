package models

import (
	"errors"
	"github.com/arangodb/go-driver"
)

type MockBooks struct{}

func (m MockBooks) Read(key string) (*Book, error) {
	if key == "correct_key" {
		return &Book{
			Key:     "correct_key",
			Title:   "some_title",
			NoPages: 42,
		}, nil
	}

	return nil, driver.ArangoError{
		Code: 404,
	}
}

func (m MockBooks) ReadAll() ([]*Book, error) {
	books := []*Book{{
		Key:     "some_key",
		Title:   "some_title",
		NoPages: 42,
	}, {
		Key:     "some_other_key",
		Title:   "some_other_title",
		NoPages: 104,
	}}

	return books, nil
}

func (m MockBooks) Delete(key string) error {
	if key != "correct_key" {
		return errors.New(NotFoundErrorCode)
	}
	return nil
}

func (m MockBooks) Create(book Book) (*Book, error) {
	if book.Key == "duplicate_key" {
		return nil, driver.ArangoError{
			Code: driver.ErrArangoConflict,
		}
	}
	return &book, nil
}

// To test errors in ReadAll use this mock
type MockBooksError struct{}

func (m MockBooksError) ReadAll() ([]*Book, error) {
	return nil, driver.ArangoError{
		Code: 500,
	}
}
func (m MockBooksError) Read(key string) (*Book, error) {
	panic("use MockBooks to test this method")
}
func (m MockBooksError) Delete(key string) error {
	panic("use MockBooks to test this method")
}
func (m MockBooksError) Create(book Book) (*Book, error) {
	panic("use MockBooks to test this method")
}
