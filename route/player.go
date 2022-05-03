package route

import (
	"cardGameSql/common"
	"cardGameSql/controller"
	"github.com/labstack/echo/v4"
)

func player(e *echo.Echo, appCtx common.AppContext) {
	player := e.Group("/api")

	var (
		authCtrl controller.Auth
	)

	player.POST("/register", authCtrl.Register(appCtx))
}