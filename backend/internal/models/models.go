package models

import (
	"github.com/arangodb/go-driver"
	"log"
)

type ModelInterface[T Book | any] interface {
	Read(key string) (*T, error)
	ReadAll() ([]T, error)
	Delete(key string) error
	Create(T) (*T, error)
}

// TODO
//type Model struct {
//	db   driver.Database
//	coll driver.Collection
//}

func New[m BookModel](dbIn driver.Database, collectionName string) *m {
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

	b := &m{
		db:   dbIn,
		coll: coll,
	}
	return b
}
