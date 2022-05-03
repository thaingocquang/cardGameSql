package controller

import (
	"cardGameSql/common"
	"cardGameSql/model"
	"cardGameSql/repository"
	"cardGameSql/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Auth struct{}

func (*Auth) Register(appCtx common.AppContext) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.PlayerRegister

		// bind data
		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// validate data
		if err := body.Validate(); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// get database from appCtx
		db := appCtx.GetDBConnection()

		// auth repo & auth service
		repo := repository.NewSQLRepo(db)
		_service := service.NewAuthService(repo)

		if err := _service.Register(&body); err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}
