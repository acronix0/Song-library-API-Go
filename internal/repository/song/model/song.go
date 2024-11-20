package model

import "time"

type Song struct {
	ID   int
	Info SongInfo
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SongInfo struct {
	ID int
	Title        string
	GroupID       int
	link 				string
	ReleaseDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Group struct{
	ID int
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}

