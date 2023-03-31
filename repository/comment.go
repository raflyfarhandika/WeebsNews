package repository

import (
	"database/sql"

	"weebsnews/database"
	"weebsnews/model"
)

type CommentRepository interface {
	Create(request model.Comment) (model.Comment, error)
	GetCommentByNewsId(request model.Comment) ([]model.Comment, error)
	GetCommentById(id int) (model.Comment, error)
	Update(request model.Comment) error
	Delete(id int) error
}

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository() CommentRepository {
	return &commentRepository{db: database.DBStart()}
}

func (comment *commentRepository) Create(request model.Comment) (model.Comment, error) {
	var result model.Comment

	statement := `INSERT INTO comment (news_id, user_id, comment) VALUES($1, $2, $3) RETURNING id, news_id, user_id, comment, created_at, updated_at`
	
	err := comment.db.QueryRow(statement, &request.NewsID, &request.UserID, &request.Comment).
		   Scan(&result.ID,
				&result.NewsID,
				&result.UserID,
				&result.Comment,
				&result.CreatedAt,
				&result.UpdatedAt)
	if err != nil {
		return model.Comment{}, err
	}

	return result, nil
}

func (comment *commentRepository) GetCommentByNewsId(request model.Comment) ([]model.Comment, error) {
	var result []model.Comment

	statement := `SELECT id,
						 news_id,
						 user_id,
						 comment,
						 created_at,
						 updated_at
				  FROM comment
				  WHERE news_id = $1`

	rows, err := comment.db.Query(statement, request.NewsID)
	if err != nil {
		return []model.Comment{}, err
	}

	for rows.Next() {
		var each model.Comment

		err = rows.Scan(&each.ID,
						&each.NewsID,
						&each.UserID,
						&each.Comment,
						&each.CreatedAt,
						&each.UpdatedAt)
		if err != nil {
			return []model.Comment{}, err
		}

		result = append(result, each)
	}

	return result, nil
}

func (comment *commentRepository) GetCommentById(id int) (model.Comment, error) {
	var result model.Comment

	statement := `SELECT id,
						 news_id,
						 user_id,
						 comment,
						 created_at,
						 updated_at
				  FROM comment
				  WHERE id = $1`

	err := comment.db.QueryRow(statement, id).
		   Scan(&result.ID,
				&result.NewsID,
				&result.UserID,
				&result.Comment,
				&result.CreatedAt,
				&result.UpdatedAt)
	if err != nil {
		return model.Comment{}, err
	}

	return result, nil
}

func (comment *commentRepository) Update(request model.Comment) error {
	statement := `UPDATE comment
				  SET comment = $1,
					  updated_at = NOW()
				  WHERE id = $2`

	err := comment.db.QueryRow(statement, request.Comment, request.ID)

	return err.Err()
}

func (comment *commentRepository) Delete(id int) error {
	statement := `DELETE FROM comment WHERE id = $1`

	err := comment.db.QueryRow(statement, id)

	return err.Err()
}