package repository

import (
	"database/sql"
	"strconv"

	"weebsnews/database"
	"weebsnews/helper"
	"weebsnews/model"
)

type UsersRepository interface {
	Create(request model.Users) (model.Users, error)
	GetAll() ([]model.Users, error)
	GetByID(id int) (model.Users, error)
	Update(request model.Users) error
	Delete(id int) error
	Login(request model.UsersLogin) (model.UsersLogin, error)
	Register(request model.Users) (model.Users, error)
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

func (users *usersRepository) GetByID(id int) (model.Users, error) {
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

func (users *usersRepository) Delete(id int) error {
	statement := "DELETE FROM users WHERE id = $1"

	err := users.db.QueryRow(statement, id)

	return err.Err()
}

func (users *usersRepository) Login(request model.UsersLogin) (model.UsersLogin, error) {
	var result model.UsersLogin

	statement := "SELECT id, role FROM users WHERE username = $1 and password = $2"

	err := users.db.QueryRow(statement, &request.Username, &request.Password).
		Scan(&result.ID, &result.Role)
	
	if err != nil {
		return model.UsersLogin{}, err
	}

	var token string

	resultID := strconv.Itoa(result.ID)

	if result != (model.UsersLogin{}) {
		token, err = helper.GetToken(resultID)

		if err != nil {
			return result, err
		}
	}

	result.Token = token

	return result, nil
}

func (users *usersRepository) Register(request model.Users) (model.Users, error) {
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