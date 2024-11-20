package song

import (
	"context"
	"database/sql"

	"github.com/acronix0/song-libary-api/internal/dto"
)



type SongRepo struct {
	db *sql.DB
}

func NewSongRepository(db *sql.DB) *SongRepo {
  return &SongRepo{db: db}
}
 
func (r *SongRepo) Create(ctx context.Context, groupName, songName string) (int, error){
	query := "INSERT INTO songs (group_name, song_name) VALUES ($1, $2)"
  _, err := r.db.ExecContext(ctx, query, groupName, songName)
  if err!= nil {
    return 1, err
  }
  return 1, nil
}

func (r *SongRepo) Get(ctx context.Context, skip, take int) ([]dto.SongDTO, error) {
	query := "SELECT id, group_name, song_name FROM songs LIMIT $1 OFFSET $2"
  rows, err := r.db.QueryContext(ctx, query, take, skip)
  if err!= nil {
    return nil, err
  }
  defer rows.Close()

  var songs []dto.SongDTO
  for rows.Next() {
    var song dto.SongDTO
    err := rows.Scan(&song.ID, &song.GroupName, &song.SongName)
    if err!= nil {
      return nil, err
    }
    songs = append(songs, song)
  }

  return songs, nil
}

func (r *SongRepo) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM songs WHERE id=$1"
  _, err := r.db.ExecContext(ctx, query, id)
  if err!= nil {
    return err
  }
  return nil
}

func (r *SongRepo) Update(ctx context.Context,  song dto.SongDTO) error {
	query := "UPDATE songs SET group_name=$1, song_name=$2 WHERE id=$3"
  _, err := r.db.ExecContext(ctx, query, song.GroupName, song.SongName, song.SongID)
  if err!= nil {
    return err
  }
  return nil
}