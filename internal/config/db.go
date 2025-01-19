package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	username := "lawrient"
	password := ""
	host := "localhost"
	port := 5432
	database := "money_management"
	maxIdleConnection := 10
	maxOpenConns := 100
	connMaxIdleTime := 5  // in minute
	connMaxLifetime := 60 // in minute

	var dsn string
	if password == "" {
		// Jika password kosong, tidak sertakan parameter password
		dsn = fmt.Sprintf(
			"host=%s user=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Makassar",
			host, username, database, port, "disable",
		)
	} else {
		// Jika password ada, sertakan parameter password
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Makassar",
			host, username, password, database, port, "disable",
		)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(maxIdleConnection)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxIdleTime(time.Duration(connMaxIdleTime) * time.Minute)
	db.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Minute)

	return db
}
