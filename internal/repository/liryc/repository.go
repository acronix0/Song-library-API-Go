package liryc

import (
	"context"
	"database/sql"

	"github.com/acronix0/song-libary-api/internal/dto"
)

type LirycsRepo struct {
	db *sql.DB
}

func NewLyricsRepository(db *sql.DB) *LirycsRepo {
  return &LirycsRepo{db: db}
}

func (r *LirycsRepo) Get(ctx context.Context,songId, skip, take int) (string, error){
	query := "SELECT id, group_name, song_name, lyrics FROM songs LIMIT $1 OFFSET $2"
  rows, err := r.db.QueryContext(ctx, query, take, skip)
  if err!= nil {
    return nil, err
  }
  defer rows.Close()
}

func (r *LirycsRepo) Create(ctx context.Context, songId int, text string) (error) {
}

func (r *LirycsRepo) Update(ctx context.Context, songId int, text string) (error) {
}

func (r *LirycsRepo) Delete(ctx context.Context, songID int) (error) {
}
