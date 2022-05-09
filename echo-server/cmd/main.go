package main

import (
	"echo-server/pkg/handler"
	"echo-server/pkg/model"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	model.InitDB()

	handler.HandleRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
