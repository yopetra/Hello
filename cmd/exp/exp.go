package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgressConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgressConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func main() {
	cfg := PostgressConfig{
		Host:     "localhost",
		Port:     "5433",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS orders(
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		amount INT,
		descrition TEXT
	);`)

	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created.")

	name := "Jon Calhoun"
	email := "jon@calhoun.com"
	_, err = db.Exec(`
		INSERT INTO users (name, email)
		VALUES ($1, $2);`, name, email)
	if err != nil {
		panic(err)
	}
	fmt.Println("User created.")
}
