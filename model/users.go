package model

import (
	"regexp"
	"time"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	Admin Role = "admin"
	User  Role = "user"
)

type Users struct {
	ID				int       `json:"id"`
	FirstName		string 	  `json:"first_name"`
	LastName		string 	  `json:"last_name"`
	Username		string 	  `json:"username"`
	Password		string    `json:"password"`
	Email			string    `json:"email"`
	Role	        Role	  `json:"role" default:"user"`
	CreatedAt		time.Time `json:"created_at"`
	UpdatedAt		time.Time `json:"updated_at"`
}

func validateEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	match, _ := regexp.MatchString(regex, email)
	return match
}

func validateEmailUnique(email string) bool {
	return true
}

func validatePassword(password string, minLength int) bool {
	return utf8.RuneCountInString(password) >= minLength
}

func (u *Users) HashPassword() error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashedPassword)
    return nil
}

func (u *Users) ComparePassword(password string) error {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    if err != nil {
        return err
    }
    return nil
}

func (u *Users) ValidateRequest() (bool, map[string]string) {
	var response = make(map[string]string)
	var isValid = true

	if u.FirstName == "" {
		isValid = false
		response["first_name"] = "First name is required"
	}

	if u.LastName == "" {
		isValid = false
		response["last_name"] = "Last name is required"
	}

	if u.Username == "" {
		isValid = false
		response["username"] = "Username is required"
	}

	if u.Password == "" {
		isValid = false
		response["password"] = "Password is required"
	} else if !validatePassword(u.Password, 6) {
		isValid = false
		response["password"] = "Password must be at least 6 characters"
	}

	if u.Email == "" {
		isValid = false
		response["email"] = "Email is required"
	} else if !validateEmail(u.Email) {
		isValid = false
		response["email"] = "Email is not valid"
	} else if !validateEmailUnique(u.Email) {
		isValid = false
		response["email"] = "Email is already taken"
	}

	return isValid, response
}