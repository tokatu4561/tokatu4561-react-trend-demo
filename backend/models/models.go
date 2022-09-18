package models

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

const dbTimeout = time.Second * 3

var db *sql.DB

// New return Models
func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		User: User{},
	}
}

type Models struct {
	User User
}

type User struct {
	UserID    int    `json:"user_id"`
	Email     string    `json:"email"`
	UserName  string    `json:"user_name,omitempty"`
	Password  string    `json:"-"`
}

// GetByEmail returns one user by email
func (u *User) GetByEmail(email string) (*User, error) {
	db, _ := sql.Open("postgres", "user=postgres password=password host=localhost port=5432 dbname=practice sslmode=disable")
	query := `select user_id, username, email, password from users where email = $1`

	var user User
	row := db.QueryRow(query, email)
	
	err := row.Scan(
		&user.UserID,
		&user.UserName,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}


// GetOne returns one user by id
func (u *User) GetById(id int) (*User, error) {
	db, _ := sql.Open("postgres", "user=postgres password=password host=localhost port=5432 dbname=practice sslmode=disable")
	query := `select user_id, username, email, password from users where user_id = $1`

	var user User
	row := db.QueryRow(query, id)
	
	err := row.Scan(
		&user.UserID,
		&user.UserName,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}