package main

import (
	"cardGameSql/appctx"
	"cardGameSql/config"
	"cardGameSql/module/database"
	"cardGameSql/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

// init ...
func init() {
	config.Init()
}

func main() {
	envVars := config.GetEnv()

	db, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	// app context
	appCtx := appctx.NewAppContext(db)

	//echo ...
	e := echo.New()

	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTION"},
		MaxAge:           600,
		AllowCredentials: false,
	}))

	//route
	route.Route(e, appCtx)

	//start server
	e.Logger.Fatal(e.Start(envVars.AppPort))
}
