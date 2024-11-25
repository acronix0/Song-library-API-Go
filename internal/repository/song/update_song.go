package song

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/acronix0/song-libary-api/internal/dto"
)

func (r *SongRepo) Update(ctx context.Context, song dto.UpdateSongDTO) (dto.ResponseSongDTO, error) {
	if song.SongID == 0 {
		return dto.ResponseSongDTO{}, fmt.Errorf("song ID is required")
	}

	var groupID *int
	if song.Group != nil {
		id, err := r.findGroupID(ctx, *song.Group)
		if err != nil {
			return dto.ResponseSongDTO{}, err
		}
		groupID = &id
	}

	query, values, err := r.buildUpdateQuery(song, groupID)
	if err != nil {
		return dto.ResponseSongDTO{}, err
	}

	updatedSong, err := r.executeUpdateQuery(ctx, query, values)
	if err != nil {
		return dto.ResponseSongDTO{}, err
	}

	return updatedSong, nil
}

func (r *SongRepo) findGroupID(ctx context.Context, groupName string) (int, error) {
	var groupID int
	query := "SELECT id FROM groups WHERE name = $1"
	err := r.db.QueryRowContext(ctx, query, groupName).Scan(&groupID)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("group not found: %s", groupName)
		}
		return 0, fmt.Errorf("failed to find group: %w", err)
	}

	return groupID, nil
}

func (r *SongRepo) buildUpdateQuery(song dto.UpdateSongDTO, groupID *int) (string, []interface{}, error) {
	var updates []string
	var values []interface{}
	argCounter := 1

	if song.Song != nil {
		updates = append(updates, fmt.Sprintf("title = $%d", argCounter))
		values = append(values, *song.Song)
		argCounter++
	}

	if groupID != nil {
		updates = append(updates, fmt.Sprintf("group_id = $%d", argCounter))
		values = append(values, *groupID)
		argCounter++
	}

	if song.Link != nil {
		updates = append(updates, fmt.Sprintf("link = $%d", argCounter))
		values = append(values, *song.Link)
		argCounter++
	}

	if song.ReleaseDate != nil {
		updates = append(updates, fmt.Sprintf("release_date = $%d", argCounter))
		values = append(values, *song.ReleaseDate)
		argCounter++
	}

	if len(updates) == 0 {
		return "", nil, fmt.Errorf("no fields to update")
	}

	updates = append(updates, fmt.Sprintf("updated_at = $%d", argCounter))
	values = append(values, time.Now())
	argCounter++

	values = append(values, song.SongID)
	query := fmt.Sprintf(`
		UPDATE songs
		SET %s
		FROM groups
		WHERE songs.group_id = groups.id AND songs.id = $%d
		RETURNING songs.id, songs.title, groups.name, songs.link, songs.release_date, songs.created_at, songs.updated_at
	`, strings.Join(updates, ", "), argCounter)

	return query, values, nil
}

func (r *SongRepo) executeUpdateQuery(ctx context.Context, query string, values []interface{}) (dto.ResponseSongDTO, error) {
	row := r.db.QueryRowContext(ctx, query, values...)

	var updatedSong dto.ResponseSongDTO
	err := row.Scan(
		&updatedSong.SongID,
		&updatedSong.Song,
		&updatedSong.Group,
		&updatedSong.Link,
		&updatedSong.ReleaseDate,
		&updatedSong.CreatedAt,
		&updatedSong.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return dto.ResponseSongDTO{}, fmt.Errorf("song not found")
		}
		return dto.ResponseSongDTO{}, fmt.Errorf("failed to update song: %w", err)
	}

	return updatedSong, nil
}
