package models

import (
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

func (m MockBooks) ReadAll() ([]Book, error) {
	books := []Book{{
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

// TODO by doing something along those lines I could test error cases for methods that dont have a parameter
//type MockBooksError struct{}
//
//func (m MockBooksError) ReadAll() ([]models.Book, error) {
//	return nil, driver.ArangoError{
//		Code: 500,
//	}
//}

func (m MockBooks) Delete(key string) error {
	//TODO implement me
	panic("implement me")
}

func (m MockBooks) Create(book Book) (*Book, error) {
	//TODO implement me
	panic("implement me")
}
