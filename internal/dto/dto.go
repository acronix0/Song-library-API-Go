package dto

import "time"

type SongDTO struct{
	  SongID       int                `json:"song_id"`         
    Song        *string            `json:"song,omitempty"` 
    Group      *string          `json:"group,omitempty"` 
		Link *string										`json:"link,omitempty"` 
    ReleaseDate  *time.Time         `json:"release_date,omitempty"` 
		CreatedAt   *time.Time						`json:"created_at,omitempty"`
	  UpdatedAt   *time.Time						`json:"updated_at,omitempty"`
    Text        *string 							`json:"lyrics,omitempty"` 
}

type LiricsDTO struct{
	  SongID       int               
		Text        string      
}

type ErrorResponse struct {
	Error string `json:"error"`
}