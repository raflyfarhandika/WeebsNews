package repository

import (
	"database/sql"

	"weebsnews/database"
	"weebsnews/model"
)

type NewsRepository interface {
	Create(request model.News) (model.News, error)
	GetAllNews() ([]model.News, error)
	GetNewsById(id int) (model.News, error)
	Update(request model.News) error
	Delete(id int) error
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
							  VALUES ($1, $2, $3, $4) 
							  RETURNING id, 
							  			user_id,  
										title, 
										content,
										thumbnail, 
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
			 &result.Thumbnail,
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
		var thumbnail sql.NullString
		err = rows.Scan(&data.ID, 
						&data.UserID, 
						&data.Title, 
						&data.Content, 
						&thumbnail, 
						&data.CreatedAt, 
						&data.UpdatedAt)
		if err != nil {
			return []model.News{}, err
		}

		if thumbnail.Valid { // cek apakah nilai thumbnail valid atau tidak
			data.Thumbnail = thumbnail.String // jika valid, gunakan nilai thumbnail yang ditemukan
		} else {
			data.Thumbnail = "" // jika tidak valid, gunakan nilai default
		}

		result = append(result, data)
	}

	return result, nil
}

func (news *newsRepository) GetNewsById(id int) (model.News, error) {
	var result model.News
	var thumbnail sql.NullString

	statement := "SELECT * FROM news WHERE id = $1"

	err := news.db.QueryRow(statement, id).
		Scan(&result.ID, 
			 &result.UserID, 
			 &result.Title, 
			 &result.Content, 
			 &thumbnail, 
			 &result.CreatedAt, 
			 &result.UpdatedAt)

	if err != nil {
		return model.News{}, err
	}

	if thumbnail.Valid { 
		result.Thumbnail = thumbnail.String 
	} else {
		result.Thumbnail = ""
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

func (news *newsRepository) Delete(id int) error {
	statement := "DELETE FROM news WHERE id = $1"

	err := news.db.QueryRow(statement, id)

	return err.Err()
}