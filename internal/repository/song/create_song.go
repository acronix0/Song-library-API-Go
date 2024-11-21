package song

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/acronix0/song-libary-api/internal/dto"
)

func (r *SongRepo) Create(ctx context.Context, song dto.SongDTO) (int, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	groupID, err := r.findOrCreateGroup(ctx, tx, song.Group)
	if err != nil {
		return 0, err
	}

	songID, err := r.insertSong(ctx, tx, song, groupID)
	if err != nil {
		return 0, err
	}

	return songID, nil
}
func (r *SongRepo) findOrCreateGroup(ctx context.Context, tx *sql.Tx, groupName *string) (int, error) {
	if groupName == nil || *groupName == "" {
		return 0, fmt.Errorf("group name is required")
	}

	var groupID int
	queryFindGroup := `SELECT id FROM groups WHERE name = $1`
	err := tx.QueryRowContext(ctx, queryFindGroup, *groupName).Scan(&groupID)

	if err != nil {
		if err == sql.ErrNoRows {
			queryCreateGroup := `
				INSERT INTO groups (name, created_at, updated_at)
				VALUES ($1, $2, $3)
				RETURNING id
			`
			err = tx.QueryRowContext(ctx, queryCreateGroup, *groupName, time.Now(), time.Now()).Scan(&groupID)
			if err != nil {
				return 0, fmt.Errorf("failed to create group: %w", err)
			}
		} else {
			return 0, fmt.Errorf("failed to find group: %w", err)
		}
	}

	return groupID, nil
}

func (r *SongRepo) insertSong(ctx context.Context, tx *sql.Tx, song dto.SongDTO, groupID int) (int, error) {
	query := `INSERT INTO songs (group_id`
	values := []interface{}{groupID} 
	placeholders := []string{"$1"}   
	argCounter := 2 //group_id is 1

	if song.Song != nil {
		query += `, title`
		values = append(values, *song.Song)
		placeholders = append(placeholders, fmt.Sprintf("$%d", argCounter))
		argCounter++
	}

	if song.Link != nil {
		query += `, link`
		values = append(values, *song.Link)
		placeholders = append(placeholders, fmt.Sprintf("$%d", argCounter))
		argCounter++
	}

	if song.ReleaseDate != nil {
		query += `, release_date`
		values = append(values, *song.ReleaseDate)
		placeholders = append(placeholders, fmt.Sprintf("$%d", argCounter))
		argCounter++
	}

	if song.Text != nil {
		query += `, text`
		values = append(values, *song.Text)
		placeholders = append(placeholders, fmt.Sprintf("$%d", argCounter))
		argCounter++
	}

	query += `, created_at, updated_at)`
	values = append(values, time.Now(), time.Now())
	placeholders = append(placeholders, fmt.Sprintf("$%d", argCounter), fmt.Sprintf("$%d", argCounter+1))

	query += ` VALUES (` + strings.Join(placeholders, ", ") + `)`
	query += ` RETURNING id`

	var songID int
	err := tx.QueryRowContext(ctx, query, values...).Scan(&songID)
	if err != nil {
		return 0, fmt.Errorf("failed to create song: %w", err)
	}

	return songID, nil
}
