package route

import (
	"cardGameSql/common"
	"cardGameSql/config"
	"github.com/labstack/echo/v4"
)

var envVars = config.GetEnv()

// Route ...
func Route(e *echo.Echo, appCtx common.AppContext) {
	player(e, appCtx)
}
