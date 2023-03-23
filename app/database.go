package app

import (
	"database/sql"
	"github.com/ariopri/golang_restfull_api/helper"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func NewDB() *sql.DB {
	// ...
	err := godotenv.Load()
	helper.PanicIfError(err)

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
