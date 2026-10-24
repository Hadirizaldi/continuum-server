package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/Hadirizaldi/continuum-server/internal/adapter"
)


func (a *Adapter) CreateSession(ctx context.Context, s *adapter.UploadSession) error {
	query := `
			INSERT INTO upload_sessions (
					id, filename, total_size, mime_type, user_id,
					chunk_size, total_chunks, status,
					created_at, last_activity_at, expires_at
			) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`
	
	_, err := a.pool.Exec(ctx, query,
		s.ID, s.Filename, s.TotalSize, s.MimeType, s.UserID,
		s.ChunkSize, s.TotalChunks, s.Status,
		s.CreatedAt, s.LastActivityAt, s.ExpiresAt,
	)

	if err != nil {
		return fmt.Errorf("create session: %w", err)
	}

	return nil
}


func(a *Adapter) GetSession (ctx context.Context, id string) (*adapter.UploadSession, error) {
	query := `
			SELECT id, filename, total_size, mime_type, user_id,
					chunk_size, total_chunks, status,
					created_at, last_activity_at, expires_at,
					completed_at, final_path, file_url
			FROM upload_sessions
			WHERE id = $1`
	
	s := &adapter.UploadSession{}

	err := a.pool.QueryRow(ctx, query, id).Scan(
			&s.ID, &s.Filename, &s.TotalSize, &s.MimeType, &s.UserID,
			&s.ChunkSize, &s.TotalChunks, &s.Status,
			&s.CreatedAt, &s.LastActivityAt, &s.ExpiresAt,
			&s.CompletedAt, &s.FinalPath, &s.FileURL,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, adapter.ErrSessionNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get session: %w", err)
	}

	return s, nil
}


func (a *Adapter) RecordChunk(ctx context.Context, c *adapter.UploadChunk, newExpiresAt time.Time) (bool, error) {
	tx, err := a.pool.Begin(ctx)
	if err != nil {
		return false, fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	chunkQuery := `
		INSERT INTO upload_chunks (session_id, chunk_index, chunk_size, offset_start, offset_end, storage_path)
		VALUES ($1,$2,$3,$4,$5,$6)
		ON CONFLICT (session_id, chunk_index) DO NOTHING`

	tag, err := tx.Exec(ctx, chunkQuery,
		c.SessionID, c.ChunkIndex, c.ChunkSize, c.OffsetStart, c.OffsetEnd, c.StoragePath,
	)
	if err != nil {
		return false, fmt.Errorf("insert chunk: %w", err)
	}

	inserted := tag.RowsAffected() > 0

	updateQuery := `
		UPDATE upload_sessions
		SET last_activity_at = now(), expires_at = $1
		WHERE id = $2`

	_, err = tx.Exec(ctx, updateQuery, newExpiresAt, c.SessionID)
	if err != nil {
		return false, fmt.Errorf("update session activity: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return false, fmt.Errorf("commit tx: %w", err)
	}
	return inserted, nil
}


func (a *Adapter) ListChunkIndexes(ctx context.Context, sessionID string) ([]int, error) {
	query := `
		SELECT chunk_index FROM upload_chunks
		WHERE session_id = $1
		ORDER BY chunk_index`

	rows, err := a.pool.Query(ctx, query, sessionID)
	if err != nil {
		return nil, fmt.Errorf("list chunks: %w", err)
	}
	defer rows.Close()

	var indexes []int
	for rows.Next() {
		var idx int
		if err := rows.Scan(&idx); err != nil {
			return nil, fmt.Errorf("scan chunk index: %w", err)
		}
		indexes = append(indexes, idx)
	}
	return indexes, rows.Err()
}


func (a *Adapter) TransitionStatus(ctx context.Context, id string, from []adapter.SessionStatus, to adapter.SessionStatus) (bool, error) {
	query := `
		UPDATE upload_sessions
		SET status = $1
		WHERE id = $2 AND status = ANY($3)`

	tag, err := a.pool.Exec(ctx, query, to, id, from)
	if err != nil {
		return false, fmt.Errorf("transition status: %w", err)
	}
	return tag.RowsAffected() > 0, nil
}


func (a *Adapter) MarkCompleted(ctx context.Context, id string, finalPath, fileURL string, completedAt time.Time) error {
	query := `
		UPDATE upload_sessions
		SET status = $1, final_path = $2, file_url = $3, completed_at = $4
		WHERE id = $5`

	_, err := a.pool.Exec(ctx, query, adapter.StatusCompleted, finalPath, fileURL, completedAt, id)
	if err != nil {
		return fmt.Errorf("mark completed: %w", err)
	}
	return nil
}


func (a *Adapter) DeleteChunks(ctx context.Context, sessionID string) error {
	query := `DELETE FROM upload_chunks WHERE session_id = $1`

	_, err := a.pool.Exec(ctx, query, sessionID)
	if err != nil {
		return fmt.Errorf("delete chunks: %w", err)
	}
	return nil
}