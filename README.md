# RetrySpool Data Storage

Data storage backend interfaces for the RetrySpool message queue system. This package handles the storage and retrieval of message content/data.

## Overview

This package provides interfaces for storing large message data efficiently. It's designed to work with backends like filesystem, S3, Azure Blob Storage, etc.

## Installation

```bash
go get schneider.vip/retryspool/storage/data
```

## Interfaces

### Backend

The core data storage interface for message content:

```go
type Backend interface {
    StoreData(ctx context.Context, messageID string, data io.Reader) (int64, error)
    GetDataReader(ctx context.Context, messageID string) (io.ReadCloser, error)
    GetDataWriter(ctx context.Context, messageID string) (io.WriteCloser, error)
    DeleteData(ctx context.Context, messageID string) error
    Close() error
}
```


### Factory

Factory pattern for creating data storage backends:

```go
type Factory interface {
    Create() (Backend, error)
    Name() string
}
```

## Usage

### Basic Usage

```go
import "schneider.vip/retryspool/storage/data"

// Create a factory (implementation-specific)
factory := filesystem.NewFactory("/path/to/data")

// Create backend
backend, err := factory.Create()
if err != nil {
    panic(err)
}
defer backend.Close()

// Store data
size, err := backend.StoreData(ctx, "msg-123", dataReader)
if err != nil {
    panic(err)
}

// Read data
reader, err := backend.GetDataReader(ctx, "msg-123")
if err != nil {
    panic(err)
}
defer reader.Close()

data, err := io.ReadAll(reader)
```


## Design Principles

- **Separation of Concerns**: Only handles message data, not metadata
- **Streaming Support**: Efficient handling of large messages
- **Backend Agnostic**: Works with any storage backend
- **Performance**: Optimized for large file operations
- **Extensible**: Support for various storage backends

## Available Implementations

- **Filesystem**: `schneider.vip/retryspool/storage/data/filesystem`
- **S3**: (planned)
- **Azure Blob**: (planned)
- **Google Cloud Storage**: (planned)

## Performance Considerations

- Use standard `io.Reader` and `io.Writer` interfaces for streaming operations
- Consider compression for text-based messages
- Implement proper connection pooling for remote backends