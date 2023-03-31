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

func (n *News) ValidateRequest() (bool, map[string]string) {
	var response = make(map[string]string)
	var isValid = true

	if n.Title == "" {
		isValid = false
		response["title"] = "Title is required"
	}

	if n.Content == "" {
		isValid = false
		response["content"] = "Content is required"
	}

	return isValid, response
}