package datastorage

import (
	"context"
	"io"
)

// Backend represents a data storage backend for message content
type Backend interface {
	// StoreData stores message data and returns the actual size written
	StoreData(ctx context.Context, messageID string, data io.Reader) (int64, error)
	
	// GetDataReader returns a reader for message data
	GetDataReader(ctx context.Context, messageID string) (io.ReadCloser, error)
	
	// GetDataWriter returns a writer for message data
	GetDataWriter(ctx context.Context, messageID string) (io.WriteCloser, error)
	
	// DeleteData removes message data
	DeleteData(ctx context.Context, messageID string) error
	
	// Close closes the data storage backend
	Close() error
}


// Factory creates data storage backends
type Factory interface {
	// Create creates a new data storage backend
	Create() (Backend, error)
	
	// Name returns the factory name
	Name() string
}