package common

import "database/sql"

type AppContext interface {
	GetDBConnection() *sql.DB
}
