package internal

import (
	"context"
	"github.com/arangodb/go-driver"
)

type MockBooks struct{}
type CollectionDocuments interface {
	ReadDocument(ctx context.Context, key string, result interface{}) (driver.DocumentMeta, error)
	CreateDocument(ctx context.Context, document interface{}) (driver.DocumentMeta, error)
}

func (c *MockBooks) ReadDocument(ctx context.Context, key string, result interface{}) (driver.DocumentMeta, error) {
	if key == "correct_key" {
		result = "test"
		return driver.DocumentMeta{}, nil
	}

	return driver.DocumentMeta{}, driver.ArangoError{
		Code: 404,
	}
}

func (c *MockBooks) RemoveDocument(ctx context.Context, key string) (driver.DocumentMeta, error) {
	if key == "correct_key" {
		return driver.DocumentMeta{}, nil
	}

	return driver.DocumentMeta{}, driver.ArangoError{
		Code: 404,
	}
}

// TODO: refactor without having to implement all those methods
func (c *MockBooks) Name() string {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) Database() driver.Database {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) Status(ctx context.Context) (driver.CollectionStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) Count(ctx context.Context) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) Statistics(ctx context.Context) (driver.CollectionStatistics, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) Revision(ctx context.Context) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) Checksum(ctx context.Context, withRevisions bool, withData bool) (driver.CollectionChecksum, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) Properties(ctx context.Context) (driver.CollectionProperties, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) SetProperties(ctx context.Context, options driver.SetCollectionPropertiesOptions) error {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) Shards(ctx context.Context, details bool) (driver.CollectionShards, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) Load(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) Unload(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) Remove(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) Truncate(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) Index(ctx context.Context, name string) (driver.Index, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) IndexExists(ctx context.Context, name string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) Indexes(ctx context.Context) ([]driver.Index, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) EnsureFullTextIndex(ctx context.Context, fields []string, options *driver.EnsureFullTextIndexOptions) (driver.Index, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) EnsureGeoIndex(ctx context.Context, fields []string, options *driver.EnsureGeoIndexOptions) (driver.Index, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) EnsureHashIndex(ctx context.Context, fields []string, options *driver.EnsureHashIndexOptions) (driver.Index, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) EnsurePersistentIndex(ctx context.Context, fields []string, options *driver.EnsurePersistentIndexOptions) (driver.Index, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) EnsureSkipListIndex(ctx context.Context, fields []string, options *driver.EnsureSkipListIndexOptions) (driver.Index, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) EnsureTTLIndex(ctx context.Context, field string, expireAfter int, options *driver.EnsureTTLIndexOptions) (driver.Index, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) EnsureZKDIndex(ctx context.Context, fields []string, options *driver.EnsureZKDIndexOptions) (driver.Index, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) EnsureInvertedIndex(ctx context.Context, options *driver.InvertedIndexOptions) (driver.Index, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) DocumentExists(ctx context.Context, key string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) ReadDocuments(ctx context.Context, keys []string, results interface{}) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) CreateDocuments(ctx context.Context, documents interface{}) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) UpdateDocument(ctx context.Context, key string, update interface{}) (driver.DocumentMeta, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) UpdateDocuments(ctx context.Context, keys []string, updates interface{}) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) ReplaceDocument(ctx context.Context, key string, document interface{}) (driver.DocumentMeta, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) ReplaceDocuments(ctx context.Context, keys []string, documents interface{}) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) RemoveDocuments(ctx context.Context, keys []string) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) ImportDocuments(ctx context.Context, documents interface{}, options *driver.ImportDocumentOptions) (driver.ImportDocumentStatistics, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockBooks) CreateDocument(ctx context.Context, document interface{}) (driver.DocumentMeta, error) {
	//TODO implement me
	panic("implement me")
}
