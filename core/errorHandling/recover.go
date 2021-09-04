package errorHandling

import (
	"github.com/labstack/echo/v4"
)

func Recover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			err := recover()
			if err != nil {
				c.Response().Header().Set("Content-Type", "application/json")

				if appErr, ok := err.(*AppError); ok {
					_ = c.JSON(appErr.Code, appErr)
					panic(err)
					return
				}

				appErr := ErrInternal(err.(error))
				_ = c.JSON(appErr.Code, appErr)
				panic(err)
				return
			}

		}()
		return next(c)
	}
}
