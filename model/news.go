package model

import "time"

type News struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Thumbnail      string    `json:"thumbnail"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}