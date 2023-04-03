package internal

import (
	"github.com/arangodb/go-driver"
	"github.com/wederer/go-chi-demo/internal/models"
)

type MockBooks2 struct{}

func (m MockBooks2) Read(key string) (*models.Book, error) {
	if key == "correct_key" {
		return &models.Book{
			Key:     "correct_key",
			Title:   "some_title",
			NoPages: 42,
		}, nil
	}

	return nil, driver.ArangoError{
		Code: 404,
	}
}

func (m MockBooks2) ReadAll() ([]models.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBooks2) Delete(key string) error {
	//TODO implement me
	panic("implement me")
}

func (m MockBooks2) Create(book models.Book) (models.Book, error) {
	//TODO implement me
	panic("implement me")
}
