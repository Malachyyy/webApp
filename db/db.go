package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const (
	host     = "192.168.2.83"
	port     = 5432
	user     = "postgres"
	password = "Welkom01"
	dbname   = "postgres"
)

func ConnectToDatabase() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	testDB(db)

	return db, nil
}

// testDB tries to ping the database
func testDB(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func AddUser(username string, password string, email string) {
	db, err := ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	hashedPassword := HashPassword(password)

	insertDynStmt := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3)`
	_, err = db.Exec(insertDynStmt, username, hashedPassword, email)
	if err != nil {
		log.Fatal(err)
	}
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}
