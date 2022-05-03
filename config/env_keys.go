package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

// env ...
var env ENV

// projectDirName ...
const projectDirName = "cardGameSql"

// InitDotEnv ...
func initDotEnv() {
	// load env
	if err := godotenv.Load(getEnvPath()); err != nil {
		log.Fatal("Error loading .env file")
	}

	// database ...
	database := Database{
		Host:     getEnvString("DB_HOST"),
		Port:     getEnvString("DB_PORT"),
		User:     getEnvString("DB_USER"),
		Password: getEnvString("DB_PASSWORD"),
		DbName:   getEnvString("DB_NAME"),
	}

	// appPort ...
	appPort := getEnvString("APP_PORT")

	// Jwt
	jwt := Jwt{SecretKey: getEnvString("SECRET_KEY")}

	env = ENV{
		Database: database,
		AppPort:  appPort,
		Jwt:      jwt,
	}

}

// getEnvPath ...
func getEnvPath() string {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	curWorkDir, _ := os.Getwd()
	rootPath := projectName.Find([]byte(curWorkDir))
	return string(rootPath) + `/.env`
}

// getEnvString ...
func getEnvString(key string) string {
	return os.Getenv(key)
}

// GetEnv ...
func GetEnv() *ENV {
	return &env
}
