package api

import (
	"link-test/api/account"

	"github.com/labstack/echo"
)

func RegisterPath(e *echo.Echo, accController *account.Controller) {
	if accController == nil {
		panic("Controller parameter cannot be nil")
	}

	acc := e.Group("/account")
	acc.GET("/:number", accController.FindBalanceByAccNo)
	acc.POST("/:from_number/transfer", accController.TransBalance)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
