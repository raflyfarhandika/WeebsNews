package repository

import (
	"database/sql"

	"weebsnews/database"
	"weebsnews/model"
)

type NewsRepository interface {
	Create(request model.News) (model.News, error)
	GetAllNews() ([]model.News, error)
	GetNewsById(id int64) (model.News, error)
	Update(request model.News) error
	Delete(request model.News) error
}

type newsRepository struct {
	db *sql.DB
}

func NewNewsRepository() NewsRepository {
	return &newsRepository{db: database.DBStart()}
}

func (news *newsRepository) Create(request model.News) (model.News, error) {
	var result model.News

	statement := `INSERT INTO news (user_id, 
									title, 
									content, 
									thumbnail) 
							  VALUES ($1, $2, $3, $4, $5) 
							  RETURNING id, 
							  			user_id,  
										title, 
										content, 
										created_at, 
										updated_at`

	err := news.db.QueryRow(statement, 
							request.UserID, 
							request.Title, 
							request.Content, 
							request.Thumbnail).
		Scan(&result.ID, 
			 &result.UserID,
			 &result.Title, 
			 &result.Content, 
			 &result.CreatedAt, 
			 &result.UpdatedAt)

	if err != nil {
		return model.News{}, err
	}

	return result, nil
}

func (news *newsRepository) GetAllNews () ([]model.News, error) {
	var result []model.News

	statement := "SELECT * FROM news"

	rows, err := news.db.Query(statement)
	if err != nil {
		return []model.News{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var data model.News
		err = rows.Scan(&data.ID, 
						&data.UserID, 
						&data.Title, 
						&data.Content, 
						&data.Thumbnail, 
						&data.CreatedAt, 
						&data.UpdatedAt)
		if err != nil {
			return []model.News{}, err
		}
		result = append(result, data)
	}

	return result, nil
}

func (news *newsRepository) GetNewsById(id int64) (model.News, error) {
	var result model.News

	statement := "SELECT * FROM news WHERE id = $1"

	err := news.db.QueryRow(statement, id).
		Scan(&result.ID, 
			 &result.UserID, 
			 &result.Title, 
			 &result.Content, 
			 &result.Thumbnail, 
			 &result.CreatedAt, 
			 &result.UpdatedAt)

	if err != nil {
		return model.News{}, err
	}

	return result, nil
}

func (news *newsRepository) Update(request model.News) error {
	statement := `UPDATE news SET title = $1, 
								  content = $2, 
								  thumbnail = $3, 
								  updated_at = NOW() 
							  WHERE id = $4`

	err := news.db.QueryRow(statement, request.Title, request.Content, request.Thumbnail, request.ID)

	return err.Err()
}

func (news *newsRepository) Delete(request model.News) error {
	statement := "DELETE FROM news WHERE id = $1"

	err := news.db.QueryRow(statement, request.ID)

	return err.Err()
}