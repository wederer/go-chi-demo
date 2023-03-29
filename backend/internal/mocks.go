package internal

import (
	"context"
	"github.com/arangodb/go-driver"
)

type BooksCollection struct{}

type CollectionDocuments interface {
	ReadDocument(ctx context.Context, key string, result interface{}) (driver.DocumentMeta, error)
	CreateDocument(ctx context.Context, document interface{}) (driver.DocumentMeta, error)
}

func (c *BooksCollection) ReadDocument(ctx context.Context, key string, result interface{}) (driver.DocumentMeta, error) {
	if key == "correct_key" {
		result = "test"
		return driver.DocumentMeta{}, nil
	} else {
		return driver.DocumentMeta{}, driver.ArangoError{
			Code: 404,
		}
	}
}

// TODO: refactor without having to implement all those methods
func (c *BooksCollection) Name() string {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) Database() driver.Database {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) Status(ctx context.Context) (driver.CollectionStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) Count(ctx context.Context) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) Statistics(ctx context.Context) (driver.CollectionStatistics, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) Revision(ctx context.Context) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) Checksum(ctx context.Context, withRevisions bool, withData bool) (driver.CollectionChecksum, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) Properties(ctx context.Context) (driver.CollectionProperties, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) SetProperties(ctx context.Context, options driver.SetCollectionPropertiesOptions) error {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) Shards(ctx context.Context, details bool) (driver.CollectionShards, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) Load(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) Unload(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) Remove(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) Truncate(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) Index(ctx context.Context, name string) (driver.Index, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) IndexExists(ctx context.Context, name string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) Indexes(ctx context.Context) ([]driver.Index, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) EnsureFullTextIndex(ctx context.Context, fields []string, options *driver.EnsureFullTextIndexOptions) (driver.Index, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) EnsureGeoIndex(ctx context.Context, fields []string, options *driver.EnsureGeoIndexOptions) (driver.Index, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) EnsureHashIndex(ctx context.Context, fields []string, options *driver.EnsureHashIndexOptions) (driver.Index, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) EnsurePersistentIndex(ctx context.Context, fields []string, options *driver.EnsurePersistentIndexOptions) (driver.Index, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) EnsureSkipListIndex(ctx context.Context, fields []string, options *driver.EnsureSkipListIndexOptions) (driver.Index, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) EnsureTTLIndex(ctx context.Context, field string, expireAfter int, options *driver.EnsureTTLIndexOptions) (driver.Index, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) EnsureZKDIndex(ctx context.Context, fields []string, options *driver.EnsureZKDIndexOptions) (driver.Index, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) EnsureInvertedIndex(ctx context.Context, options *driver.InvertedIndexOptions) (driver.Index, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) DocumentExists(ctx context.Context, key string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) ReadDocuments(ctx context.Context, keys []string, results interface{}) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) CreateDocuments(ctx context.Context, documents interface{}) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) UpdateDocument(ctx context.Context, key string, update interface{}) (driver.DocumentMeta, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) UpdateDocuments(ctx context.Context, keys []string, updates interface{}) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) ReplaceDocument(ctx context.Context, key string, document interface{}) (driver.DocumentMeta, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) ReplaceDocuments(ctx context.Context, keys []string, documents interface{}) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) RemoveDocument(ctx context.Context, key string) (driver.DocumentMeta, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) RemoveDocuments(ctx context.Context, keys []string) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) ImportDocuments(ctx context.Context, documents interface{}, options *driver.ImportDocumentOptions) (driver.ImportDocumentStatistics, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BooksCollection) CreateDocument(ctx context.Context, document interface{}) (driver.DocumentMeta, error) {
	//TODO implement me
	panic("implement me")
}
