package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	
	db, err := sql.Open("postgres", "user=postgres password=password host=localhost port=5432 dbname=practice sslmode=disable")
    if err != nil {
        log.Fatalln("接続失敗", err)
    }
    defer db.Close()

	err = createUser(db)
	if err != nil {
		log.Fatalln(err)
	}
}

// AddUser inserts a user into the database
func createUser(sql *sql.DB) error {
	stmt := `
		insert into users (username, password, email)
		values ($1, $2, $3)`

	newHash, err := bcrypt.GenerateFromPassword([]byte("12341234"), 12)
	if err != nil {
		return err
	}

	_, err = sql.Exec(stmt,
		"テストユーザー",
		string(newHash),
		"test@example.com",
	)
	if err != nil {
		return err
	}

	return nil
}