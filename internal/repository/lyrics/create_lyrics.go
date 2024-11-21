package lyrics

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func (r *LyricsRepo) Create(ctx context.Context, songID int, text string) error {
	if songID == 0 || text == "" {
		return fmt.Errorf("song ID and text are required")
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	verses := strings.Split(text, "\n")

	query := `
		INSERT INTO song_lyrics (song_id, verse_number, text, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	for i, verse := range verses {
		_, err = tx.ExecContext(ctx, query, songID, i+1, verse, time.Now(), time.Now())
		if err != nil {
			return fmt.Errorf("failed to insert verse %d: %w", i+1, err)
		}
	}

	return nil
}
