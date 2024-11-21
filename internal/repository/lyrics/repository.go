package lyrics

import (
	"database/sql"
)

type LyricsRepo struct {
	db *sql.DB
}

func NewLyricsRepository(db *sql.DB) *LyricsRepo {
  return &LyricsRepo{db: db}
}