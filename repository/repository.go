package repository

import "database/sql"

// sqlRepo ...
type sqlRepo struct {
	db *sql.DB
}

// NewSQLRepo ...
func NewSQLRepo(db *sql.DB) *sqlRepo {
	return &sqlRepo{db: db}
}
