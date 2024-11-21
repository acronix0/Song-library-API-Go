package lyrics

import (
	"context"
	"fmt"
)

func (r *LyricsRepo) Delete(ctx context.Context, songID int) error {
	if songID == 0 {
		return fmt.Errorf("song ID is required")
	}

	query := `DELETE FROM song_lyrics WHERE song_id = $1`
	res, err := r.db.ExecContext(ctx, query, songID)
	if err != nil {
		return fmt.Errorf("failed to delete lyrics: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no lyrics found for song ID %d", songID)
	}

	return nil
}