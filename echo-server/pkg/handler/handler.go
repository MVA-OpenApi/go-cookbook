package handler

import "github.com/labstack/echo/v4"

func HandleRoutes(e *echo.Echo) {
	mainStoreRoutes(e)
}
