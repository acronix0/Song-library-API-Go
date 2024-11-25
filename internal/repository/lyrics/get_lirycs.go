package lyrics

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func (r *LyricsRepo) Get(ctx context.Context, songID, skip, take int) (string, error) {
	query := `
		SELECT text
		FROM song_lyrics
		WHERE song_id = $1
		ORDER BY verse_number ASC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, songID, take, skip)
	if err != nil {
		return "", fmt.Errorf("failed to fetch lyrics: %w", err)
	}
	defer rows.Close()

	var lyrics []string
	for rows.Next() {
		var verse string
		if err := rows.Scan(&verse); err != nil {
			return "", fmt.Errorf("failed to scan row: %w", err) // Обработка ошибки `rows.Scan`
		}
		lyrics = append(lyrics, verse)
	}

	if err := rows.Err(); err != nil {
		return "", fmt.Errorf("error iterating over rows: %w", err)
	}

	if len(lyrics) == 0 {
		return "", sql.ErrNoRows
	}

	return strings.Join(lyrics, "\n"), nil
}
