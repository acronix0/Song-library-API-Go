package song

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/acronix0/song-libary-api/internal/dto"
)

func (r *SongRepo) Get(ctx context.Context, skip, take int) ([]dto.ResponseSongDTO, error) {
	query := `
		SELECT 
			s.id AS song_id,
			s.title,
			g.name AS group_name,
			s.link,
			s.release_date,
			s.created_at,
			s.updated_at,
			COALESCE(STRING_AGG(sl.text, E'\n' ORDER BY sl.verse_number), '') AS lyrics
		FROM songs s
		LEFT JOIN groups g ON s.group_id = g.id
		LEFT JOIN song_lyrics sl ON sl.song_id = s.id
		GROUP BY s.id, g.name
		ORDER BY s.created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, take, skip)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch songs: %w", err)
	}
	defer rows.Close()

	var songs []dto.ResponseSongDTO
	for rows.Next() {
		var song dto.ResponseSongDTO

		err := rows.Scan(
			&song.SongID,
			&song.Song,
			&song.Group,
			&song.Link,
			&song.ReleaseDate,
			&song.CreatedAt,
			&song.UpdatedAt,
			&song.Text,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		songs = append(songs, song)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", rows.Err())
	}
	if songs == nil {
		return []dto.ResponseSongDTO{}, sql.ErrNoRows
	}
	return songs, nil
}
