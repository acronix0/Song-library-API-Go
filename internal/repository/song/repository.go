package song

import (
	"database/sql"
)



type SongRepo struct {
	db *sql.DB
}

func NewSongRepository(db *sql.DB) *SongRepo {
  return &SongRepo{db: db}
}
 

