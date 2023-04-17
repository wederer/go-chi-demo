package models

type ModelInterface[T any] interface {
	Read(key string) (*T, error)
	ReadAll() ([]*T, error)
	Delete(key string) error
	Create(T) (*T, error)
}
