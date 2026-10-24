package adapter

import (
	"context"
	"errors"
	"io"
	"time"
)


var (
	ErrSessionNotFound = errors.New("session not found")
	ErrStorageFull = errors.New("storage full")
	ErrChunkNotFound = errors.New("chunk not found")
)

type SessionStatus string

const (
	StatusPending		SessionStatus = "pending"
	StatusUploading		SessionStatus = "uploading"
	StatusMerging		SessionStatus = "merging"
	StatusCompleted		SessionStatus = "completed"
	StatusFailed		SessionStatus = "failed"
	StatusCancelled		SessionStatus = "cancelled"
)

type UploadSession struct {
	ID             string
	Filename       string
	TotalSize      int64
	MimeType       string
	UserID         *string
	ChunkSize      int64
	TotalChunks    int
	Status         SessionStatus
	CreatedAt      time.Time
	LastActivityAt time.Time
	ExpiresAt      time.Time
	CompletedAt    *time.Time
	FinalPath      *string
	FileURL        *string
}

type UploadChunk struct {
	SessionID   string
	ChunkIndex  int
	ChunkSize   int64
	OffsetStart int64
	OffsetEnd   int64
	StoragePath string
}


type DatabaseAdapter interface {
	CreateSession(ctx context.Context, s *UploadSession) error
	GetSession(ctx context.Context, id string) (*UploadSession, error)
	RecordChunk(ctx context.Context, c *UploadChunk, newExpiresAt time.Time) (inserted bool, err error)
	ListChunkIndexes(ctx context.Context, sessionID string) ([]int, error)
	TransitionStatus(ctx context.Context, id string, from []SessionStatus, to SessionStatus) (ok bool, err error)
	MarkCompleted(ctx context.Context, id string, finalPath, fileURL string, completedAt time.Time) error
	DeleteChunks(ctx context.Context, sessionID string) error
}


type StorageAdapter interface {
	PrepareSession(ctx context.Context, sessionID string) error
	SaveChunk(ctx context.Context, sessionID string, chunkIndex int, r io.Reader, size int64) (storagePath string, err error)
	ChunkExists(ctx context.Context, sessionID string, chunkIndex int) (bool, error)
	MergeChunks(ctx context.Context, sessionID string, totalChunks int, expectedSize int64, finalPath string) (actualFinalPath string, err error)
	DeleteSession(ctx context.Context, sessionID string) error
}

