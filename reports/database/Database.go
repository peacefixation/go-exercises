package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Postgres Driver
)

// DB the database connection
type DB struct {
	conn *sql.DB
	tx   *sql.Tx
}

// New instance of the database connection
func New(credentials Credentials) (*DB, error) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=require",
		credentials.User,
		credentials.Password,
		credentials.Name,
	)

	connection, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, fmt.Errorf("Failed to initialise database connection. %v", err)
	}

	err = connection.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed to ping database connection. %v", err)
	}

	return &DB{conn: connection}, nil
}

// Query execute a query that returns rows
func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if db.tx == nil {
		return db.conn.Query(query, args...)
	}
	return db.tx.Query(query, args...)
}

// QueryRow execute a query that returns a single row
func (db *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	if db.tx == nil {
		return db.conn.QueryRow(query, args...)
	}
	return db.tx.QueryRow(query, args...)
}

// Exec execute a query that doesn't return rows
func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	if db.tx == nil {
		return db.conn.Exec(query, args...)
	}
	return db.tx.Exec(query, args...)
}

// Begin a database transaction if one doesn't exist
func (db *DB) Begin() (bool, error) {
	if db.tx != nil {
		return false, nil
	}

	tx, err := db.conn.Begin()
	if err != nil {
		return false, err
	}

	db.tx = tx
	return true, err
}

// Rollback if a database transaction exists
func (db *DB) Rollback() error {
	if db.tx == nil {
		return nil
	}

	db.tx = nil
	return db.tx.Rollback()
}

// Commit if a database transaction exists
func (db *DB) Commit() error {
	if db.tx == nil {
		return nil
	}

	db.tx = nil
	return db.tx.Commit()
}
