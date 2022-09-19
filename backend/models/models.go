package models

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const dbTimeout = time.Second * 3

var db *sql.DB

type Models struct {
	User User
}

type User struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	UserName string `json:"user_name,omitempty"`
	Password string `json:"-"`
}

type Task struct {
	ID        int       `db:"id" json:"id"`
	UserID    int       `db:"user_id" json:"user_id"`
	Title     string    `db:"title" json:"title"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// New return Models
func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		User: User{},
	}
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

func (t *Task) GetById(id int) (*Task, error) {
	db, _ := sql.Open("postgres", "user=postgres password=password host=localhost port=5432 dbname=practice sslmode=disable")
	query := `select id, user_id, title, created_at, updated_at from tasks where id = $1`

	var task Task
	row := db.QueryRow(query, id)

	err := row.Scan(
		&task.ID,
		&task.UserID,
		&task.Title,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *Task) Insert(user_id int, task *Task) (*Task, error) {
	stmt := `insert into tasks (user_id, title, created_at, updated_at)
		values ($1, $2, $3, $4) returning id`

	result, err := db.Exec(stmt,
		user_id,
		task.Title,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}
	log.Println(result)

	return nil, nil
}

func (t *Task) Update(task *Task) error {
	stmt := `
			update tasks set
			title = $1
			updated_at = $2
			where id = $3
			`

	_, err := db.Exec(stmt,
		task.Title,
		time.Now(),
		task.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (t *Task) Delete(task_id int) error {
	stmt := `
			delete from tasks where id = $1
			`

	_, err := db.Exec(stmt, task_id)
	if err != nil {
		return err
	}

	return nil
}
