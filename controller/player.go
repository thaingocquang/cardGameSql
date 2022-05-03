package controller

import (
	"cardGameSql/common"
	"cardGameSql/repository"
	"cardGameSql/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// GetAllPlayer ...
func GetAllPlayer(appCtx common.AppContext) echo.HandlerFunc {
	return func(c echo.Context) error {
		db := appCtx.GetDBConnection()

		repo := repository.NewSQLRepo(db)
		playerService := service.NewPlayerService(repo)

		players, err := playerService.GetAllPlayer()
		if err != nil {
			fmt.Println(err)
			return err
		}

		return c.JSON(http.StatusOK, players)

	}
}

// GetPlayerByID ...
func GetPlayerByID(appCtx common.AppContext) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idNum, _ := strconv.Atoi(id)

		db := appCtx.GetDBConnection()
		repo := repository.NewSQLRepo(db)
		playerService := service.NewPlayerService(repo)

		player, err := playerService.GetPlayerByID(idNum)
		if err != nil {
			fmt.Println(err)
			return err
		}

		return c.JSON(http.StatusOK, player)

	}
}
