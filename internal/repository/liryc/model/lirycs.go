package model

import "time"

type SongLyrics struct {
	ID          int
	SongID      int
	VerseNumber int
	Text        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}