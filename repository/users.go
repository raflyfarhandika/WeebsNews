package repository

import (
	"database/sql"

	"weebsnews/database"
	"weebsnews/model"
)

type UsersRepository interface {
	Create(request model.Users) (model.Users, error)
	GetAll() ([]model.Users, error)
	GetByID(id int64) (model.Users, error)
	Update(request model.Users) error
	Delete(request model.Users) error
}

type usersRepository struct {
	db *sql.DB
}

func NewUsersRepository() UsersRepository {
	return &usersRepository{db: database.DBStart()}
}

func (users *usersRepository) Create(request model.Users) (model.Users, error) {
	var result model.Users

	statement := `INSERT INTO users (first_name, last_name, username, password, email, role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, first_name, last_name, username, password, email, role, created_at, updated_at`

	request.HashPassword()
	request.Role = "user"

	err := users.db.QueryRow(statement, &request.FirstName, &request.LastName, &request.Username, &request.Password, &request.Email, &request.Role).
		Scan(&result.ID, &result.FirstName, &result.LastName, &result.Username, &result.Password, &result.Email, &result.Role, &result.CreatedAt, &result.UpdatedAt)

	if err != nil {
		return model.Users{}, err
	}

	return result, nil
}

func (users *usersRepository) GetAll() ([]model.Users, error) {
	var result []model.Users

	statement := "SELECT id, first_name, last_name, username, email, role, created_at, updated_at FROM users"

	rows, err := users.db.Query(statement)
	if err != nil {
		return []model.Users{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var data model.Users
		err = rows.Scan(&data.ID, &data.FirstName, &data.LastName, &data.Username, &data.Email, &data.Role, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return []model.Users{}, err
		}
		result = append(result, data)
	}

	return result, nil
}

func (users *usersRepository) GetByID(id int64) (model.Users, error) {
	var result model.Users

	statement := "SELECT id, first_name, last_name, username, email, role, created_at, updated_at FROM users WHERE id = $1"

	err := users.db.QueryRow(statement, id).
		Scan(&result.ID, &result.FirstName, &result.LastName, &result.Username, &result.Email, &result.Role, &result.CreatedAt, &result.UpdatedAt)
	
	if err != nil {
		return model.Users{}, err
	}

	return result, nil
}

func (users *usersRepository) Update(request model.Users) error {
	statement := "UPDATE users SET first_name = $1, last_name = $2, email = $3, updated_at = NOW() WHERE id = $4"

	err := users.db.QueryRow(statement, &request.FirstName, &request.LastName, &request.Email, &request.ID)

	return err.Err()
}

func (users *usersRepository) Delete(request model.Users) error {
	statement := "DELETE FROM users WHERE id = $1"

	err := users.db.QueryRow(statement, &request.ID)

	return err.Err()
}