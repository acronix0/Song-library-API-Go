package song

import (
	"context"
	"fmt"

	"github.com/acronix0/song-libary-api/internal/dto"
)

func (r *SongRepo) Get(ctx context.Context, skip, take int) ([]dto.SongDTO, error) {
	query := `
		SELECT 
			s.id AS song_id,
			s.title,
			g.name AS group_name,
			s.link,
			s.release_date,
			s.created_at,
			s.updated_at,
			s.text
		FROM songs s
		LEFT JOIN groups g ON s.group_id = g.id
		ORDER BY s.created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, take, skip)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch songs: %w", err)
	}
	defer rows.Close()

	var songs []dto.SongDTO
	for rows.Next() {
		var song dto.SongDTO
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

	return songs, nil
}
