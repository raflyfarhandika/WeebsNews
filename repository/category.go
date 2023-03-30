package repository

import (
	"database/sql"

	"weebsnews/database"
	"weebsnews/model"
)

type CategoryRepository interface {
	Create(request model.Category) (model.Category, error)
	GetAll() ([]model.Category, error)
	GetByID(request model.Category) (model.Category, error)
	Update(request model.Category) error
	Delete(request model.Category) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{db: database.DBStart()}
}

func (category *categoryRepository) Create(request model.Category) (model.Category, error) {
	var result model.Category

	statement := "INSERT INTO category (category_name) VALUES ($1) RETURNING id, category_name, created_at, updated_at"

	err := category.db.QueryRow(statement, &request.CategoryName).
		Scan(&result.ID, &result.CategoryName, &result.CreatedAt, &result.UpdatedAt)

	if err != nil {
		return model.Category{}, err
	}
	
	return result, nil

}

func (category *categoryRepository) GetAll() ([]model.Category, error) {
	var result []model.Category

	statement := "SELECT * FROM category"

	rows, err := category.db.Query(statement)
	if err != nil {
		return []model.Category{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var data model.Category
		err = rows.Scan(&data.ID, &data.CategoryName, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return []model.Category{}, err
		}
		result = append(result, data)
	}

	return result, nil
}

func (category *categoryRepository) GetByID(request model.Category) (model.Category, error) {
	var result model.Category

	statement := "SELECT * FROM category WHERE id = $1"

	err := category.db.QueryRow(statement, request.ID).
		Scan(&result.ID, &result.CategoryName, &result.CreatedAt, &result.UpdatedAt)

	if err != nil {
		return model.Category{}, err
	}

	return result, nil
}

func (category *categoryRepository) Update(request model.Category) error {
	statement := "UPDATE category SET category_name = $1 WHERE id = $2"

	err := category.db.QueryRow(statement, &request.CategoryName, &request.ID)

	return err.Err()
}

func (category *categoryRepository) Delete(request model.Category) error {
	statement := "DELETE FROM category WHERE id = $1"

	err := category.db.QueryRow(statement, &request.ID)

	return err.Err()
}