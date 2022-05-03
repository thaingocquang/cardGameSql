package middleware

import (
	"cardGameSql/common"
	"fmt"
	"github.com/labstack/echo/v4"
)

// Recover ...
func Recover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				// if AppError error
				if appErr, ok := err.(*common.AppError); ok {
					// abort request and return json
					c.JSON(appErr.StatusCode, appErr)
				}
				// if system error
				appErr := common.ErrInternal(err.(error))
				fmt.Println(appErr)
			}
		}()
		return next(c)
	}
}
