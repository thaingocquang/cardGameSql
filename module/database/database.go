package database

import (
	"cardGameSql/config"
	"database/sql"
	"fmt"
)

func Connect() (*sql.DB, error) {
	envVars := config.GetEnv()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		envVars.Database.Host, envVars.Database.Port, envVars.Database.User, envVars.Database.Password, envVars.Database.DbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Database successfully connected!")
	return db, nil
}
