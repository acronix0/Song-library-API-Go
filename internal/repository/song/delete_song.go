package song

import (
	"context"
	"fmt"
)

func (r *SongRepo) Delete(ctx context.Context, songID int) error {
	if songID == 0 {
		return fmt.Errorf("song ID is required")
	}

	query := `DELETE FROM songs WHERE id = $1`
	res, err := r.db.ExecContext(ctx, query, songID)
	if err != nil {
		return fmt.Errorf("failed to delete song: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("song with ID %d not found", songID)
	}

	return nil
}