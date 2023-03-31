package model

import "time"

type Comment struct {
	ID        int       `json:"id"`
	NewsID    int       `json:"news_id"`
	UserID    int       `json:"user_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}