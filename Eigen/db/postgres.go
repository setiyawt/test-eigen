package db

import (
	"database/sql"
	"fmt"
	"myproject/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type Postgres struct{}

func Connect(creds *model.Credential) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func SQLExecute(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, code VARCHAR(255), name VARCHAR(255), password VARCHAR(255))")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS sessions (id SERIAL PRIMARY KEY, token VARCHAR(255), name VARCHAR(255), expiry timestamp default NULL)")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS books (id SERIAL PRIMARY KEY, code VARCHAR(6), title VARCHAR(255), author VARCHAR(255), Stock INT)")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS borrowed (id SERIAL PRIMARY KEY, code_book VARCHAR(5), code_member VARCHAR(4), borrowedDate timestamp default NULL, returnedDate timestamp default NULL, Status VARCHAR(10), late INT, quantity INT)")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS penalties (id SERIAL PRIMARY KEY, code_member VARCHAR(4), penalty_type VARCHAR(50), penalty_amount DECIMAL(10, 2), penalty_active BOOLEAN, penalty_date TIMESTAMP, resolved_date TIMESTAMP)")
	if err != nil {
		return err
	}

	return nil
}

func Reset(db *sql.DB, table string) error {
	_, err := db.Exec("TRUNCATE " + table)
	if err != nil {
		return err
	}

	_, err = db.Exec("ALTER SEQUENCE " + table + "_id_seq RESTART WITH 1")
	if err != nil {
		return err
	}

	return nil
}
