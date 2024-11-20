package dto

import "time"

type SongDTO struct{
	  SongID       int                `json:"song_id"`         
    Title        *string            `json:"title,omitempty"` 
    GroupName      *string          `json:"group_id,omitempty"` 
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