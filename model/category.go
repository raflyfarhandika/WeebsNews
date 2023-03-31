package model

import "time"

type Category struct {
	ID				int       `json:"id"`
	CategoryName	string    `json:"category_name"`
	CreatedAt		time.Time `json:"created_at"`
	UpdatedAt		time.Time `json:"updated_at"`
}

type CategoryRelation struct {
	ID				int       `json:"id"`
	NewsID			int       `json:"news_id"`
	CategoryID		int       `json:"category_id"`
	CreatedAt		time.Time `json:"created_at"`
	UpdatedAt		time.Time `json:"updated_at"`
}

func (c *Category) ValidateRequest() (bool, map[string]string) {
	var response = make(map[string]string)
	var isValid = true

	if c.CategoryName == "" {
		isValid = false
		response["category_name"] = "Category name is required"
	}

	return isValid, response
}